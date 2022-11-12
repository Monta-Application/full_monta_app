# Generated by Django 4.1.2 on 2022-10-30 23:40

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0004_alter_organization_language_and_more"),
        ("integration", "0004_alter_integration_auth_token"),
    ]

    operations = [
        migrations.AlterField(
            model_name="integration",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
    ]
