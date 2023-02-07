# Generated by Django 4.1.5 on 2023-01-09 08:19

import encrypted_model_fields.fields
from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("worker", "0003_alter_workercomment_worker"),
    ]

    operations = [
        migrations.AlterField(
            model_name="workerprofile",
            name="license_number",
            field=encrypted_model_fields.fields.EncryptedCharField(
                blank=True,
                help_text="Driver License Number",
                verbose_name="License Number",
            ),
        ),
    ]
