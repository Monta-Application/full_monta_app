# Generated by Django 4.1.5 on 2023-01-12 06:15

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("stops", "0002_alter_stop_location"),
    ]

    operations = [
        migrations.AlterField(
            model_name="stop",
            name="address_line",
            field=models.CharField(
                blank=True,
                help_text="Stop Address",
                max_length=255,
                verbose_name="Stop Address",
            ),
        ),
    ]
