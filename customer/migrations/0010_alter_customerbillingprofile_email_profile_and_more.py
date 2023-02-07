# Generated by Django 4.1.5 on 2023-02-02 03:22

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("customer", "0009_alter_customerbillingprofile_customer_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="customerbillingprofile",
            name="email_profile",
            field=models.ForeignKey(
                default=1,
                help_text="Customer Email Profile",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="billing_profile",
                to="customer.customeremailprofile",
                verbose_name="Customer Email Profile",
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="customerbillingprofile",
            name="rule_profile",
            field=models.ForeignKey(
                default=1,
                help_text="Rule Profile",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="billing_profile",
                to="customer.customerruleprofile",
                verbose_name="Rule Profile",
            ),
            preserve_default=False,
        ),
    ]
