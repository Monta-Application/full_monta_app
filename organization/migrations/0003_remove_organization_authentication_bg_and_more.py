# Generated by Django 4.1.4 on 2023-01-05 03:14

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0002_alter_depotdetail_depot"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="organization",
            name="authentication_bg",
        ),
        migrations.RemoveField(
            model_name="organization",
            name="authentication_template",
        ),
    ]
