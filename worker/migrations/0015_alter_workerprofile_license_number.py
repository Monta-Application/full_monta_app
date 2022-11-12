# Generated by Django 4.1.2 on 2022-11-05 19:52

from django.db import migrations, models

import dispatch.validators.regulatory


class Migration(migrations.Migration):

    dependencies = [
        ("worker", "0014_alter_workerprofile_endorsements_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="workerprofile",
            name="license_number",
            field=models.CharField(
                blank=True,
                help_text="License Number.",
                max_length=20,
                null=True,
                validators=[
                    dispatch.validators.regulatory.validate_worker_regulatory_information
                ],
                verbose_name="License Number",
            ),
        ),
    ]
