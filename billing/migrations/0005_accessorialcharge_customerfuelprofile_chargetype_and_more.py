# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-11-07 05:43

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0006_alter_depotdetail_address_line_1_and_more"),
        ("order", "0005_ordercontrol"),
        ("billing", "0004_customerfueltable_customerfueltabledetail"),
    ]

    operations = [
        migrations.CreateModel(
            name="AccessorialCharge",
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
                    "code",
                    models.CharField(
                        max_length=50,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "is_fuel_surcharge",
                    models.BooleanField(
                        default=False, verbose_name="Is Fuel Surcharge"
                    ),
                ),
                (
                    "is_detention",
                    models.BooleanField(default=False, verbose_name="Is Detention"),
                ),
                (
                    "method",
                    models.CharField(
                        choices=[("D", "Distance"), ("F", "Flat"), ("P", "Percentage")],
                        default="D",
                        max_length=1,
                        verbose_name="Method",
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
                "verbose_name": "Other Charge",
                "verbose_name_plural": "Other Charges",
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
                    models.DateField(help_text="End Date", verbose_name="End Date"),
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
                        help_text="Base Price",
                        max_digits=16,
                        null=True,
                        verbose_name="Base Price",
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
                        to="billing.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "customer_fuel_table",
                    models.ForeignKey(
                        blank=True,
                        help_text="Customer Fuel Profile",
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="billing.customerfueltable",
                        verbose_name="Customer Fuel Profile",
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
            name="ChargeType",
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
                    "name",
                    models.CharField(max_length=50, unique=True, verbose_name="Name"),
                ),
                (
                    "description",
                    models.CharField(
                        blank=True,
                        max_length=100,
                        null=True,
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
                "verbose_name": "Charge Type",
                "verbose_name_plural": "Charge Types",
            },
        ),
        migrations.AddIndex(
            model_name="chargetype",
            index=models.Index(fields=["name"], name="billing_cha_name_b28a59_idx"),
        ),
    ]
