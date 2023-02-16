# Generated by Django 4.1.7 on 2023-02-16 02:17

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0008_alter_tablechangealert_options_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="tablechangealert",
            name="function_name",
            field=models.CharField(
                default=1,
                help_text="The function name that the table change alert will use.",
                max_length=50,
                verbose_name="Function Name",
            ),
            preserve_default=False,
        ),
        migrations.AddField(
            model_name="tablechangealert",
            name="trigger_name",
            field=models.CharField(
                default=1,
                help_text="The trigger name that the table change alert will use.",
                max_length=50,
                verbose_name="Trigger Name",
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="tablechangealert",
            name="database_action",
            field=models.CharField(
                choices=[("INSERT", "Insert"), ("UPDATE", "Update"), ("BOTH", "Both")],
                default="INSERT",
                help_text="The database action that the table change alert is for.",
                max_length=50,
                verbose_name="Database Action",
            ),
        ),
    ]
