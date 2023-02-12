# Generated by Django 4.1.6 on 2023-02-12 03:09

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("invoicing", "0002_invoicecontrol_attach_pdf_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="invoicecontrol",
            name="invoice_logo_width",
            field=models.PositiveIntegerField(
                default=0,
                help_text="Define invoice logo width. (PX)",
                verbose_name="Invoice Logo Width",
            ),
        ),
        migrations.AddField(
            model_name="invoicecontrol",
            name="show_invoice_due_date",
            field=models.BooleanField(
                default=True,
                help_text="Show the invoice due date on the invoice.",
                verbose_name="Show Invoice Due Date",
            ),
        ),
    ]
