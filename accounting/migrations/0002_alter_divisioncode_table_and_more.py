# Generated by Django 4.1.7 on 2023-02-18 16:30

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("accounting", "0001_initial"),
    ]

    operations = [
        migrations.AlterModelTable(
            name="divisioncode",
            table="division_code",
        ),
        migrations.AlterModelTable(
            name="generalledgeraccount",
            table="general_ledger_account",
        ),
        migrations.AlterModelTable(
            name="revenuecode",
            table="revenue_code",
        ),
    ]
