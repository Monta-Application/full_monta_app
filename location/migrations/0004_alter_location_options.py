# Generated by Django 4.1.5 on 2023-01-19 17:16

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("location", "0003_alter_locationcontact_location"),
    ]

    operations = [
        migrations.AlterModelOptions(
            name="location",
            options={
                "ordering": ("code",),
                "verbose_name": "Location",
                "verbose_name_plural": "Locations",
            },
        ),
    ]
