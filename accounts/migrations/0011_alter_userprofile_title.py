# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-11-07 18:31

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("accounts", "0010_remove_user_organization"),
    ]

    operations = [
        migrations.AlterField(
            model_name="userprofile",
            name="title",
            field=models.ForeignKey(
                blank=True,
                null=True,
                on_delete=django.db.models.deletion.PROTECT,
                related_name="profile",
                related_query_name="profiles",
                to="accounts.jobtitle",
                verbose_name="Job Title",
            ),
        ),
    ]
