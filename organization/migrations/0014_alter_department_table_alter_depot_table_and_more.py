# Generated by Django 4.1.7 on 2023-02-18 16:24

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0013_alter_tablechangealert_options_and_more"),
    ]

    operations = [
        migrations.AlterModelTable(
            name="department",
            table="department",
        ),
        migrations.AlterModelTable(
            name="depot",
            table="depot",
        ),
        migrations.AlterModelTable(
            name="depotdetail",
            table="depot_detail",
        ),
        migrations.AlterModelTable(
            name="emailcontrol",
            table="email_control",
        ),
        migrations.AlterModelTable(
            name="emaillog",
            table="email_log",
        ),
        migrations.AlterModelTable(
            name="emailprofile",
            table="email_profile",
        ),
        migrations.AlterModelTable(
            name="organization",
            table="organization",
        ),
        migrations.AlterModelTable(
            name="tablechangealert",
            table="table_change_alert",
        ),
        migrations.AlterModelTable(
            name="taxrate",
            table="tax_rate",
        ),
    ]
