# Generated by Django 4.1.2 on 2022-10-30 02:30

from django.db import migrations, models

from organization.validators.organization import validate_org_timezone


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="organization",
            name="timezone",
            field=models.CharField(
                default="America/New_York",
                help_text="The timezone of the organization",
                max_length=255,
                validators=[validate_org_timezone],
                verbose_name="Timezone",
            ),
        ),
    ]
