# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------
import json
import uuid
from typing import TYPE_CHECKING, Any

import redis
from django.core.exceptions import ValidationError
from django.db import IntegrityError, transaction
from django.db.models import QuerySet
from django.shortcuts import get_object_or_404
from django.utils import timezone

from accounts.models import User
from billing import exceptions, models, selectors, utils
from order.models import Order
from utils.services.pdf import UUIDEncoder

if TYPE_CHECKING:
    from utils.types import ModelUUID


def generate_invoice_number(
    *, instance: models.BillingQueue, is_credit_memo: bool = False
) -> str:
    """
    Generate an invoice number for a given BillingQueue instance.

    This function generates an invoice number for a given BillingQueue instance, taking into
    account whether the instance is a credit memo or not. For non-credit memos, the function
    increments the invoice number with a letter suffix (A, B, C, etc.) for each new invoice
    related to the same order. For credit memos, it assigns the invoice number of the latest
    invoice related to the same order.

    Args:
        instance (models.BillingQueue): The instance for which to generate an invoice number.
        is_credit_memo (bool, optional): Flag indicating whether the instance is a credit memo.
            Defaults to False.

    Returns:
        str: The generated invoice number.

    Time Complexity: O(1) - Constant time is required to generate the invoice number,
        assuming that database lookups take constant time.
    """
    prefix = instance.organization.invoice_control.invoice_number_prefix
    order_pro_number = instance.order.pro_number
    if instance.order.billing_queue.exists():
        latest_invoice = instance.order.billing_queue.latest("created")

        if is_credit_memo:
            instance.invoice_number = latest_invoice.invoice_number
        else:
            suffixes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

            if latest_invoice.invoice_number[-1] in suffixes:
                next_suffix = suffixes[
                    suffixes.index(latest_invoice.invoice_number[-1]) + 1
                ]
                instance.invoice_number = (
                    latest_invoice.invoice_number[:-1] + next_suffix
                )
            else:
                instance.invoice_number = f"{latest_invoice.invoice_number}A"
    else:
        instance.invoice_number = f"{prefix}{order_pro_number}"

    return instance.invoice_number


@transaction.atomic
def transfer_to_billing_queue_service(
    *, user_id: "ModelUUID", order_pros: list[str], task_id: str
) -> str:
    """
    Atomically transfers eligible orders to the billing queue, logs the transfer,
    and returns a success message. If any part of the operation fails, all changes are rolled back.

    Args:
        user_id (ModelUUID): The ID of the user transferring the orders.
        order_pros (List[str]): A list of order PRO numbers to be transferred.
        task_id (str): The ID of the task that initiated the transfer.

    Returns:
        str: A message indicating the success of the transfer and the number of orders transferred.

    Raises:
        exceptions.BillingException: If no eligible orders are found for transfer or if an error occurs
            while transferring an order. In case of an error, the transaction is aborted, ensuring that
            no orders are transferred if there's a problem with any of them.

    Time Complexity: O(n), where n is the number of orders. The main operations (creating BillingQueue
        objects, updating Order objects, and creating BillingTransferLog objects) are performed for each order.
        However, these operations are managed efficiently using bulk operations.
    """
    # Get the user
    user = get_object_or_404(User, id=user_id)

    billing_control = user.organization.billing_control

    # Get the billable orders
    orders = selectors.get_billable_orders(
        organization=user.organization, order_pros=order_pros
    )

    # If there are no orders, raise an BillingException
    if not orders:
        raise exceptions.BillingException(
            f"No orders found to be eligible for transfer. Orders must be marked {billing_control.order_transfer_criteria}"
        )

    # Get the current time
    now = timezone.now()

    # Create a list of BillingTransferLog objects
    transfer_log = []

    # Loop through the orders and create a BillingQueue object for each
    for order in orders:
        try:
            # Create a BillingQueue object
            models.BillingQueue.objects.create(
                organization=order.organization,
                order=order,
                customer=order.customer,
                business_unit=order.business_unit,
            )

            # Update the order
            order.transferred_to_billing = True
            order.billing_transfer_date = now

            # Create a BillingTransferLog object
            transfer_log.append(
                models.BillingTransferLog(
                    order=order,
                    organization=order.organization,
                    business_unit=order.business_unit,
                    task_id=task_id,
                    transferred_at=now,
                    transferred_by=user,
                )
            )

        # If there is a ValidationError or IntegrityError, create a BillingException
        except (ValidationError, IntegrityError) as error:
            error_type = type(error).__name__
            utils.create_billing_exception(
                user=user,
                exception_type="OTHER",
                invoice=order,
                exception_message=f"Order {order.pro_number} failed to transfer to billing queue: {error_type} - {error}",
            )

    # Bulk update the orders
    Order.objects.bulk_update(
        orders, ["transferred_to_billing", "billing_transfer_date"]
    )

    # Bulk create the transfer log
    models.BillingTransferLog.objects.bulk_create(transfer_log)

    # Return a success message
    return f"Successfully transferred {len(orders)} orders to billing queue."


def mass_order_billing_service(
    *, user_id: "ModelUUID", task_id: str | uuid.UUID
) -> None:
    """Process the billing for multiple orders.

    Args:
        user_id (ModelUUID): The ID of the user initiating the mass billing.
        task_id (str | uuid.UUID): The ID of the task that initiated the mass billing.

    Returns:
        None: This function does not return anything.
    """

    user: User = get_object_or_404(User, id=user_id)
    orders = selectors.get_billing_queue(user=user, task_id=task_id)
    bill_orders(user_id=user_id, invoices=orders)


def bill_orders(
    *,
    user_id: "ModelUUID",
    invoices: QuerySet[models.BillingQueue] | models.BillingQueue,
) -> tuple[list[str], list[str]] | None:
    """
    Bills the specified orders. If the organization enforces customer billing requirements,
    checks these requirements before billing the order. If requirements are not met, a
    BillingException is created.

    Args:
        user_id (ModelUUID): The ID of the user responsible for billing the orders.
        invoices (QuerySet[models.BillingQueue] | models.BillingQueue): An iterable of BillingQueue instances
            or a single BillingQueue instance representing the orders to be billed.

    Returns:
        None: This function does not return anything.

    Raises:
        Http404: If the user with the given user_id does not exist.
        exceptions.BillingException: If the customer billing requirements are not met.

    Space Complexity: O(n), where n is the number of invoices. This is because a list of invoices
        is created in memory when a single BillingQueue instance is provided. However, the function
        does not create additional data structures that grow with the size of the input.

    Time Complexity: O(n), where n is the number of invoices. The function performs operations (checking
        billing requirements and calling 'order_billing_actions') for each invoice. The actual time complexity
        might be affected by these operations and how they are implemented.
    """
    user = get_object_or_404(User, id=user_id)
    billed_invoices = []
    order_missing_info = []

    # If invoices is a BillingQueue object, convert it to a list
    if isinstance(invoices, models.BillingQueue):
        invoices = [invoices]  # type: ignore

    # Check the organization enforces customer billing_requirements
    organization_enforces_billing = utils.check_organization_enforces_customer_billing(
        organization=user.organization
    )

    # Loop through the invoices and bill them
    for invoice in invoices:
        if organization_enforces_billing:
            # If the organization enforces customer billing requirements, check the requirements
            _, missing_documents = utils.check_billing_requirements(
                user=user, invoice=invoice
            )
            # Append missing_documents only when it is not empty
            if missing_documents:
                order_missing_info.append(missing_documents)
            else:
                # If the customer billing requirements are met, bill the order
                utils.order_billing_actions(invoice=invoice, user=user)
                billed_invoices.append(invoice.invoice_number)

        else:
            # If the customer billing requirements are met or not enforced, bill the order
            utils.order_billing_actions(invoice=invoice, user=user)
            billed_invoices.append(invoice.invoice_number)
    return order_missing_info, billed_invoices


def untransfer_order_service(invoice_numbers: QuerySet[models.BillingQueue]) -> None:
    """Untransfer the specified orders from the billing queue.

    Args:
        invoice_numbers (QuerySet[models.BillingQueue]): QuerySet of BillingQueue objects to be untransferred.

    Returns:
        None: This function does not return anything.
    """

    for invoice_number in invoice_numbers:
        invoice_number.order.transferred_to_billing = False
        invoice_number.order.billing_transfer_date = None
        invoice_number.order.save()
        invoice_number.delete()


def ready_to_bill_service(order: QuerySet[Order]) -> None:
    """Automatically set orders ready to bill, if order passes billing requirement check.

    Args:
        order (QuerySet[Order]): Order Queryset

    Returns:
        None: This function does not return anything.
    """
    for order in order:
        organization = order.organization

        if organization.billing_control.auto_mark_ready_to_bill:
            if utils.check_billing_requirements(user=order.created_by, invoice=order):
                order.ready_to_bill = True
                order.save()
        elif order.customer.auto_mark_ready_to_bill:
            if utils.check_billing_requirements(user=order.created_by, invoice=order):
                order.ready_to_bill = True
                order.save()


class BillingClientSessionManager:
    """
    Manages client sessions for billing through a Redis datastore.

    Attributes:
        client (redis.Redis): Redis client used for session management.

    Args:
        host (str): The hostname of the Redis server. Defaults to "localhost".
        port (int): The port of the Redis server. Defaults to 6379.
        db (int): The DB number to connect to on the Redis server. Defaults to 4.
    """

    def __init__(self, host: str = "localhost", port: int = 6379, db: int = 4):
        """
        Constructs all the necessary attributes for the BillingClientSessionManager object.

        Args:
            host (str): The hostname of the Redis server. Defaults to "localhost".
            port (int): The port of the Redis server. Defaults to 6379.
            db (int): The DB number to connect to on the Redis server. Defaults to 4.
        """
        self.client = redis.Redis(host=host, port=port, db=db)

    @staticmethod
    def _get_session_key(user_id: uuid.UUID) -> str:
        """
        Generates a session key string based on user_id.

        Args:
            user_id (uuid.UUID): The unique identifier for the user.

        Returns:
            str: A session key string.
        """
        return f"billing_client_session_{user_id}"

    @staticmethod
    def _serialize(data: Any) -> str:
        """
        Serializes the given data into a JSON string.

        Args:
            data (dict): The data to be serialized.

        Returns:
            str: The serialized data.
        """
        return json.dumps(data, cls=UUIDEncoder)

    @staticmethod
    def _deserialize(data: Any) -> dict[str, Any]:
        """
        Deserializes the given JSON string into a Python object.
        If data is None, returns None.

        Args:
            data (str): The JSON string to be deserialized.

        Returns:
            dict: The deserialized data if data is not None.
            None: If data is None.
        """
        return json.loads(data)

    def set_new_billing_client_session(self, user_id: uuid.UUID):
        """
        Sets a new billing client session for a user.

        If a session already exists, this function fetches and returns it.
        Otherwise, a new session is created with a predefined structure and stored in Redis.

        Args:
            user_id (uuid.UUID): The unique identifier for the user.

        Returns:
            dict: The client session.
        """
        session_key = self._get_session_key(user_id)
        user_sessions = self.client.get(session_key)
        if user_sessions is None:
            billing_client_session = {
                "user_id": user_id,
                "current_action": "getting_started",
                "previous_action": None,
                "last_response": None,
                "last_payload": None,
            }
            self.client.set(session_key, self._serialize(billing_client_session))
        else:
            billing_client_session = self._deserialize(user_sessions)
        return billing_client_session

    def update_billing_client_session(
        self, user_id: uuid.UUID, billing_client_session: dict[str, Any]
    ) -> None:
        """
        Updates a user's billing client session in the Redis datastore.

        Args:
            user_id (uuid.UUID): The unique identifier for the user.
            billing_client_session (dict): The updated session data.

        Returns:
            None: This function does not return anything.
        """
        session_key = self._get_session_key(user_id)
        self.client.set(session_key, self._serialize(billing_client_session))

    def get_billing_client_session(self, user_id: uuid.UUID) -> dict[str, Any] | None:
        """
        Retrieves a user's billing client session from the Redis datastore.

        Args:
            user_id (uuid.UUID): The unique identifier for the user.

        Returns:
            dict: The client session if it exists.
            None: If no session is found for the user.
        """
        session_key = self._get_session_key(user_id)
        billing_client_session = self.client.get(session_key)
        return self._deserialize(billing_client_session)

    def delete_billing_client_session(self, user_id: uuid.UUID) -> None:
        """
        Deletes a user's billing client session from the Redis datastore.

        Args:
            user_id (uuid.UUID): The unique identifier for the user.

        Returns:
            None: This function does not return anything.
        """
        session_key = self._get_session_key(user_id)
        self.client.delete(session_key)
