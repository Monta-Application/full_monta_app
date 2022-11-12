# Generated by Django 4.1.2 on 2022-11-08 18:25

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import localflavor.us.models
import phonenumber_field.modelfields


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("billing", "0012_invoicetemplate_invoicetype_and_more"),
        ("order", "0006_alter_commodity_description_and_more"),
    ]

    operations = [
        migrations.CreateModel(
            name="Customer",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer should be treated as active. Unselect this instead of deleting customers.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "code",
                    models.CharField(
                        editable=False,
                        help_text="Customer code",
                        max_length=10,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Customer name", max_length=150, verbose_name="Name"
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        help_text="Address line 1",
                        max_length=150,
                        verbose_name="Address Line 1",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="Address line 2",
                        max_length=150,
                        verbose_name="Address Line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="City", max_length=150, verbose_name="City"
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        help_text="State", max_length=2, verbose_name="State"
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        help_text="Zip code", max_length=10, verbose_name="Zip Code"
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer",
                "verbose_name_plural": "Customers",
                "ordering": ["code"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelTable",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.CharField(
                        editable=False,
                        help_text="Customer Fuel Profile ID",
                        max_length=10,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="ID",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        help_text="Customer Fuel Profile Description",
                        max_length=150,
                        verbose_name="Description",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Fuel Table",
                "verbose_name_plural": "Customer Fuel Table",
                "ordering": ["id"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelTableDetail",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "amount",
                    models.DecimalField(
                        blank=True,
                        decimal_places=5,
                        help_text="Amount",
                        max_digits=16,
                        null=True,
                        verbose_name="Amount",
                    ),
                ),
                (
                    "method",
                    models.CharField(
                        choices=[("D", "Distance"), ("F", "Flat"), ("P", "Percentage")],
                        help_text="Method",
                        max_length=1,
                        verbose_name="Method",
                    ),
                ),
                (
                    "start_price",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Start Price",
                        max_digits=5,
                        null=True,
                        verbose_name="Start Price",
                    ),
                ),
                (
                    "percentage",
                    models.DecimalField(
                        blank=True,
                        decimal_places=2,
                        help_text="Percentage",
                        max_digits=6,
                        null=True,
                        verbose_name="Percentage",
                    ),
                ),
                (
                    "customer_fuel_table",
                    models.ForeignKey(
                        help_text="Customer Fuel Profile",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_table_details",
                        related_query_name="customer_fuel_table_detail",
                        to="customer.customerfueltable",
                        verbose_name="Customer Fuel Profile",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Fuel Profile Detail",
                "verbose_name_plural": "Customer Fuel Profile Details",
                "ordering": ["customer_fuel_table"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelProfile",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "start_date",
                    models.DateField(help_text="Start Date", verbose_name="Start Date"),
                ),
                (
                    "end_date",
                    models.DateField(
                        blank=True,
                        help_text="End Date",
                        null=True,
                        verbose_name="End Date",
                    ),
                ),
                (
                    "days_to_use",
                    models.CharField(
                        choices=[
                            ("D", "Delivery Date"),
                            ("S", "Actual Shipment Date"),
                            ("C", "Scheduled Shipment Date"),
                            ("E", "Date Entered"),
                        ],
                        help_text="Days to Use",
                        max_length=1,
                        verbose_name="Days to Use",
                    ),
                ),
                (
                    "fuel_region",
                    models.CharField(
                        choices=[
                            ("USA", "US National Average"),
                            ("EAST", "East Coast"),
                            ("NE", "New England"),
                            ("GA", "General Atlantic"),
                            ("LA", "Lower Atlantic"),
                            ("MW", "Midwest"),
                            ("GC", "Gulf Coast"),
                            ("RM", "Rocky Mountain"),
                            ("WC", "West Coast"),
                            ("CA", "California"),
                            ("WCL", "West Coast (No LA)"),
                        ],
                        help_text="Fuel Region",
                        max_length=4,
                        verbose_name="Fuel Region",
                    ),
                ),
                (
                    "fsc_method",
                    models.CharField(
                        choices=[
                            ("P", "Percentage"),
                            ("F", "Flat"),
                            ("D", "Distance"),
                            ("T", "Table"),
                        ],
                        help_text="FSC Method",
                        max_length=1,
                        verbose_name="FSC Method",
                    ),
                ),
                (
                    "base_price",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Base Price",
                        max_digits=16,
                        null=True,
                        verbose_name="Base Price",
                    ),
                ),
                (
                    "fuel_variance",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Fuel Variance ex: 0.02",
                        max_digits=16,
                        null=True,
                        verbose_name="Fuel Variance",
                    ),
                ),
                (
                    "amount",
                    models.DecimalField(
                        blank=True,
                        decimal_places=5,
                        help_text="Amount",
                        max_digits=16,
                        null=True,
                        verbose_name="Amount",
                    ),
                ),
                (
                    "percentage",
                    models.DecimalField(
                        blank=True,
                        decimal_places=2,
                        help_text="Percentage",
                        max_digits=6,
                        null=True,
                        verbose_name="Percentage",
                    ),
                ),
                (
                    "customer",
                    models.ForeignKey(
                        help_text="Customer",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="customer.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "customer_fuel_table",
                    models.ForeignKey(
                        blank=True,
                        help_text="Customer Fuel Table",
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="customer.customerfueltable",
                        verbose_name="Customer Fuel Table",
                    ),
                ),
                (
                    "fsc_code",
                    models.ForeignKey(
                        help_text="FSC Code",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="billing.accessorialcharge",
                        verbose_name="FSC Code",
                    ),
                ),
                (
                    "order_type",
                    models.ForeignKey(
                        help_text="Order Type",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="order.ordertype",
                        verbose_name="Order Type",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Fuel Profile",
                "verbose_name_plural": "Customer Fuel Profiles",
                "ordering": ["customer"],
            },
        ),
        migrations.CreateModel(
            name="CustomerContact",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer contact should be treated as active. Unselect this instead of deleting customer contacts.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Contact name",
                        max_length=150,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        blank=True,
                        help_text="Contact email",
                        max_length=150,
                        verbose_name="Email",
                    ),
                ),
                (
                    "title",
                    models.CharField(
                        blank=True,
                        help_text="Contact title",
                        max_length=100,
                        verbose_name="Title",
                    ),
                ),
                (
                    "phone",
                    phonenumber_field.modelfields.PhoneNumberField(
                        blank=True,
                        help_text="Contact phone",
                        max_length=20,
                        null=True,
                        region=None,
                        verbose_name="Phone Number",
                    ),
                ),
                (
                    "is_payable_contact",
                    models.BooleanField(
                        default=False,
                        help_text="Designates whether this contact is the payable contact",
                        verbose_name="Payable Contact",
                    ),
                ),
                (
                    "customer",
                    models.ForeignKey(
                        help_text="Customer",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="contacts",
                        related_query_name="contact",
                        to="customer.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Contact",
                "verbose_name_plural": "Customer Contacts",
                "ordering": ["customer", "name"],
            },
        ),
        migrations.CreateModel(
            name="CustomerBillingProfile",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer billing profile should be treated as active. Unselect this instead of deleting customer billing profiles.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "customer",
                    models.OneToOneField(
                        help_text="Customer",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="billing_profile",
                        related_query_name="billing_profiles",
                        to="customer.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "document_class",
                    models.ManyToManyField(
                        help_text="Document class",
                        related_name="billing_profiles",
                        related_query_name="billing_profile",
                        to="billing.documentclassification",
                        verbose_name="Document Class",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Billing Profile",
                "verbose_name_plural": "Customer Billing Profiles",
                "ordering": ["customer"],
            },
        ),
    ]
