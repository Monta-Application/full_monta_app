# Generated by Django 4.1.3 on 2022-11-19 00:48

from django.db import migrations

from utils import models


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0006_dispatchcontrol_generate_routes"),
    ]

    operations = [
        migrations.AlterField(
            model_name="dispatchcontrol",
            name="distance_method",
            field=models.ChoiceField(
                choices=[("Google", "Google"), ("Monta", "Monta")],
                default="Monta",
                help_text="Distance method for the company.",
                max_length=6,
                verbose_name="Distance Method",
            ),
        ),
        migrations.AlterField(
            model_name="dispatchcontrol",
            name="record_service_incident",
            field=models.ChoiceField(
                choices=[
                    ("Never", "Never"),
                    ("Pickup", "Pickup"),
                    ("Delivery", "Delivery"),
                    ("Pickup and Delivery", "Pickup and Delivery"),
                    ("All except shipper", "All except shipper"),
                ],
                default="Never",
                max_length=19,
                verbose_name="Record Service Incident",
            ),
        ),
    ]
