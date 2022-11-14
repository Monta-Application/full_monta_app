# Generated by Django 4.1.2 on 2022-11-12 18:40

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("route", "0003_alter_route_destination_alter_route_origin"),
    ]

    operations = [
        migrations.AlterField(
            model_name="routecontrol",
            name="organization",
            field=models.OneToOneField(
                help_text="Organization related to this route control",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="route_controls",
                related_query_name="route_control",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
    ]
