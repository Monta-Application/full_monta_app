"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

import pytest
from django.core.exceptions import ValidationError
from django.urls import reverse
from pydantic import BaseModel
import uuid
from accounting import models
from accounting.tests.factories import GeneralLedgerAccountFactory

pytestmark = pytest.mark.django_db


class GeneralLedgerAccountBase(BaseModel):
    """
    Division Code Base Schema
    """

    organization_id: uuid.UUID
    is_active: bool
    account_number: str
    description: str
    account_type: str
    cash_flow_type: str
    account_sub_type: str
    account_classification: str


class GeneralAccountLedgerAccountCreate(GeneralLedgerAccountBase):
    """
    Division Code Create Schema
    """

    pass


class GeneralAccountLedgerAccountUpdate(GeneralLedgerAccountBase):
    """
    Division Code Update Schema
    """

    id: uuid.UUID


def test_create_schema() -> None:
    """
    Test create schema
    """
    gl_account_create = GeneralAccountLedgerAccountCreate(
        organization_id=uuid.uuid4(),
        is_active=True,
        account_number="1234-1234-1234-1234",
        description="Description",
        account_type=models.GeneralLedgerAccount.AccountTypeChoices.ASSET,
        cash_flow_type=models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING,
        account_sub_type=models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET,
        account_classification=models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE,
    )

    gl_account = gl_account_create.dict()

    assert gl_account is not None
    assert gl_account["is_active"] is True
    assert gl_account["account_number"] == "1234-1234-1234-1234"
    assert gl_account["description"] == "Description"
    assert (
        gl_account["account_type"]
        == models.GeneralLedgerAccount.AccountTypeChoices.ASSET
    )
    assert (
        gl_account["cash_flow_type"]
        == models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING
    )
    assert (
        gl_account["account_sub_type"]
        == models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET
    )
    assert (
        gl_account["account_classification"]
        == models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE
    )


def test_update_schema() -> None:
    """
    Test General Ledger Account update Schema
    """

    gl_account_update = GeneralAccountLedgerAccountUpdate(
        id=uuid.uuid4(),
        organization_id=uuid.uuid4(),
        is_active=True,
        account_number="1234-1234-1234-1234",
        description="Description",
        account_type=models.GeneralLedgerAccount.AccountTypeChoices.ASSET,
        cash_flow_type=models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING,
        account_sub_type=models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET,
        account_classification=models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE,
    )

    gl_account = gl_account_update.dict()

    assert gl_account is not None
    assert gl_account["is_active"] is True
    assert gl_account["account_number"] == "1234-1234-1234-1234"
    assert gl_account["description"] == "Description"
    assert (
        gl_account["account_type"]
        == models.GeneralLedgerAccount.AccountTypeChoices.ASSET
    )
    assert (
        gl_account["cash_flow_type"]
        == models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING
    )
    assert (
        gl_account["account_sub_type"]
        == models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET
    )
    assert (
        gl_account["account_classification"]
        == models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE
    )


def test_delete_schema() -> None:
    """
    Test GL Account Delete Schema
    """

    gl_accounts = [
        GeneralLedgerAccountBase(
            organization_id=uuid.uuid4(),
            is_active=True,
            account_number="1234-1234-1234-1234",
            description="Description 1",
            account_type=models.GeneralLedgerAccount.AccountTypeChoices.ASSET,
            cash_flow_type=models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING,
            account_sub_type=models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET,
            account_classification=models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE,
        ),
        GeneralLedgerAccountBase(
            organization_id=uuid.uuid4(),
            is_active=True,
            account_number="1234-1234-1234-1235",
            description="Description 2",
            account_type=models.GeneralLedgerAccount.AccountTypeChoices.ASSET,
            cash_flow_type=models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING,
            account_sub_type=models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET,
            account_classification=models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE,
        ),
    ]

    # Store the GL Accounts in a list
    gl_account_store = gl_accounts.copy()

    # Delete the first GL Account
    gl_account_store.pop(0)

    assert len(gl_accounts) == 2
    assert len(gl_account_store) == 1
    assert gl_accounts[0].account_number == "1234-1234-1234-1234"
    assert gl_account_store[0].account_number == "1234-1234-1234-1235"


def test_gl_account_str_representation(general_ledger_account) -> None:
    """
    Test GL Account String Representation
    """
    assert str(general_ledger_account) == general_ledger_account.account_number


def test_gl_account_get_absolute_url(general_ledger_account) -> None:
    """
    Test GL Account Get Absolute URL
    """
    assert (
        general_ledger_account.get_absolute_url()
        == f"/api/gl_accounts/{general_ledger_account.id}/"
    )


def test_gl_account_clean_method_with_valid_data(general_ledger_account) -> None:
    """
    Test GL Account Clean Method with valid data
    """
    try:
        general_ledger_account.clean()
    except ValidationError:
        pytest.fail("clean method raised ValidationError unexpectedly")


def test_list(general_ledger_account) -> None:
    """
    Test general ledger account list
    """
    assert general_ledger_account is not None


def test_create(general_ledger_account) -> None:
    """
    Test general ledger account creation
    """
    gl_account = models.GeneralLedgerAccount.objects.create(
        organization=general_ledger_account.organization,
        account_number="1234-1234-1234-1234",
        account_type=models.GeneralLedgerAccount.AccountTypeChoices.ASSET,
        description="Another Description",
        cash_flow_type=models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING,
        account_sub_type=models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET,
        account_classification=models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_PAYABLE,
    )

    assert gl_account is not None
    assert gl_account.account_number == "1234-1234-1234-1234"


def test_update(general_ledger_account) -> None:
    """
    Test general ledger account update
    """
    general_ledger_account.account_number = "1234-1234-1234-1234"
    general_ledger_account.save()
    assert general_ledger_account.account_number == "1234-1234-1234-1234"


def test_api_get(api_client) -> None:
    """
    Test get General Ledger accounts
    """
    response = api_client.get(reverse("gl-accounts-list"))
    assert response.status_code == 200


def test_api_get_by_id(api_client, gl_account_api) -> None:
    """
    Test get General Ledger account by ID
    """
    response = api_client.get(
        reverse(
            "gl-accounts-detail",
            kwargs={"pk": gl_account_api.data["id"]},
        )
    )
    assert response.status_code == 200
    assert response.data["account_number"] == gl_account_api.data["account_number"]
    assert response.data["account_type"] == gl_account_api.data["account_type"]
    assert response.data["description"] == gl_account_api.data["description"]


def test_api_put(api_client, gl_account_api) -> None:
    """
    Test put General Ledger Account
    """
    response = api_client.put(
        reverse(
            "gl-accounts-detail",
            kwargs={"pk": gl_account_api.data["id"]},
        ),
        {
            "account_number": "2345-2345-2345-2345",
            "description": "Another Test Description",
            "account_type": f"{models.GeneralLedgerAccount.AccountTypeChoices.EXPENSE}",
            "cash_flow_type": f"{models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING}",
            "account_sub_type": f"{models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET}",
            "account_classification": f"{models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_RECEIVABLE}",
        },
    )
    assert response.status_code == 200
    assert response.data["account_number"] == "2345-2345-2345-2345"
    assert (
        response.data["account_type"]
        == models.GeneralLedgerAccount.AccountTypeChoices.EXPENSE
    )
    assert response.data["description"] == "Another Test Description"
    assert (
        response.data["cash_flow_type"]
        == models.GeneralLedgerAccount.CashFlowTypeChoices.FINANCING
    )
    assert (
        response.data["account_sub_type"]
        == models.GeneralLedgerAccount.AccountSubTypeChoices.CURRENT_ASSET
    )
    assert (
        response.data["account_classification"]
        == models.GeneralLedgerAccount.AccountClassificationChoices.ACCOUNTS_RECEIVABLE
    )


def test_api_delete(api_client, gl_account_api) -> None:
    """
    Test delete general Ledger account
    """
    response = api_client.delete(
        reverse(
            "gl-accounts-detail",
            kwargs={"pk": gl_account_api.data["id"]},
        )
    )

    assert response.status_code == 200
    assert not response.data


def test_account_number(general_ledger_account) -> None:
    """
    Test Whether the validation error is thrown if the entered account_number value is not a
    regex match.
    """

    with pytest.raises(ValidationError) as excinfo:
        general_ledger_account.account_number = "00000-2323411-124141"
        general_ledger_account.full_clean()

    assert excinfo.value.message_dict["account_number"] == [
        "Account number must be in the format 0000-0000-0000-0000."
    ]


def test_unique_account_numer(general_ledger_account) -> None:
    """
    Test creating a General Ledger account with the same account number
    throws ValidationError.
    """
    general_ledger_account.account_number = "1234-1234-1234-1234"
    general_ledger_account.save()

    with pytest.raises(ValidationError) as excinfo:
        GeneralLedgerAccountFactory(account_number="1234-1234-1234-1234")

    assert excinfo.value.message_dict["account_number"] == [
        "An account with this account number already exists. Please try again."
    ]
