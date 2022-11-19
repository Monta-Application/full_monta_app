# Generated by Django 4.1.3 on 2022-11-16 20:17

from django.db import migrations

from utils import models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0024_alter_commodity_unit_of_measure"),
    ]

    operations = [
        migrations.AlterField(
            model_name="commodity",
            name="unit_of_measure",
            field=models.ChoiceField(
                blank=True,
                choices=[
                    ("PALLET", "Pallet"),
                    ("TOTE", "Tote"),
                    ("DRUM", "Drum"),
                    ("CYLINDERRRRRRRRRR", "Cylinder"),
                    ("CASE", "Case"),
                    ("AMPULE", "Ampule"),
                    ("BAG", "Bag"),
                    ("BOTTLE", "Bottle"),
                    ("PAIL", "Pail"),
                    ("PIECES", "Pieces"),
                    ("ISO_TANK", "ISO Tank"),
                ],
                help_text="Unit of Measure of the Commodity",
                max_length=17,
                verbose_name="Unit of Measure",
            ),
        ),
    ]
