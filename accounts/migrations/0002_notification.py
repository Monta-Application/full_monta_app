# Generated by Django 4.1.6 on 2023-02-12 03:17

import django.db.models.deletion
import django.utils.timezone
import jsonfield.fields
from django.conf import settings
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0002_alter_emailprofile_protocol_taxrate"),
        ("contenttypes", "0002_remove_content_type_name"),
        ("accounts", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="Notification",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "level",
                    models.CharField(
                        choices=[
                            ("success", "success"),
                            ("info", "info"),
                            ("warning", "warning"),
                            ("error", "error"),
                        ],
                        default="info",
                        max_length=20,
                    ),
                ),
                ("unread", models.BooleanField(db_index=True, default=True)),
                ("actor_object_id", models.CharField(max_length=255)),
                ("verb", models.CharField(max_length=255)),
                ("description", models.TextField(blank=True, null=True)),
                (
                    "target_object_id",
                    models.CharField(blank=True, max_length=255, null=True),
                ),
                (
                    "action_object_object_id",
                    models.CharField(blank=True, max_length=255, null=True),
                ),
                (
                    "timestamp",
                    models.DateTimeField(
                        db_index=True, default=django.utils.timezone.now
                    ),
                ),
                ("public", models.BooleanField(db_index=True, default=True)),
                ("deleted", models.BooleanField(db_index=True, default=False)),
                ("emailed", models.BooleanField(db_index=True, default=False)),
                ("data", jsonfield.fields.JSONField(blank=True, null=True)),
                (
                    "action_object_content_type",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="notify_action_object",
                        to="contenttypes.contenttype",
                    ),
                ),
                (
                    "actor_content_type",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="notify_actor",
                        to="contenttypes.contenttype",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="notifications",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "recipient",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="notifications",
                        to=settings.AUTH_USER_MODEL,
                    ),
                ),
                (
                    "target_content_type",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="notify_target",
                        to="contenttypes.contenttype",
                    ),
                ),
            ],
            options={
                "verbose_name": "Notification",
                "verbose_name_plural": "Notifications",
            },
        ),
    ]
