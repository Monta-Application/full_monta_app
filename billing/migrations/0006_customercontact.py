# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-11-07 20:51

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import phonenumber_field.modelfields


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0006_alter_depotdetail_address_line_1_and_more"),
        ("billing", "0005_accessorialcharge_customerfuelprofile_chargetype_and_more"),
    ]

    operations = [
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
                        help_text="Contact name", max_length=150, verbose_name="Name"
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        help_text="Contact email", max_length=150, verbose_name="Email"
                    ),
                ),
                (
                    "title",
                    models.CharField(
                        help_text="Contact title", max_length=100, verbose_name="Title"
                    ),
                ),
                (
                    "phone",
                    phonenumber_field.modelfields.PhoneNumberField(
                        help_text="Contact phone",
                        max_length=20,
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
                        to="billing.customer",
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
    ]
