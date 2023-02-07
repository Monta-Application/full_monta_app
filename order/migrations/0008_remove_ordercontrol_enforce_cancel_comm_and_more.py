# Generated by Django 4.1.5 on 2023-01-19 02:49

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("order", "0007_alter_orderdocumentation_document"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="ordercontrol",
            name="enforce_cancel_comm",
        ),
        migrations.RemoveField(
            model_name="ordercontrol",
            name="enforce_shipper",
        ),
        migrations.AddField(
            model_name="ordercontrol",
            name="enforce_voided_comm",
            field=models.BooleanField(
                default=False,
                help_text="Enforce comment when voiding an order.",
                verbose_name="Enforce Voided Comm",
            ),
        ),
        migrations.AlterField(
            model_name="ordercontrol",
            name="calculate_distance",
            field=models.BooleanField(
                default=True,
                help_text="Automatically Calculate distance for the order",
                verbose_name="Calculate Distance",
            ),
        ),
    ]
