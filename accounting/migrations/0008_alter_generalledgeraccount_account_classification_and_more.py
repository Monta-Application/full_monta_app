# Generated by Django 4.1.3 on 2022-11-19 00:48

from django.db import migrations

from utils import models


class Migration(migrations.Migration):

    dependencies = [
        ("accounting", "0007_remove_generalledgeraccount_account_sub_type_en_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="generalledgeraccount",
            name="account_classification",
            field=models.ChoiceField(
                blank=True,
                choices=[
                    ("BANK", "Bank"),
                    ("CASH", "Cash"),
                    ("ACCOUNTS_RECEIVABLE", "Accounts Receivable"),
                    ("ACCOUNTS_PAYABLE", "Accounts Payable"),
                    ("INVENTORY", "Inventory"),
                    ("OTHER_CURRENT_ASSET", "Other Current Asset"),
                    ("FIXED_ASSET", "Fixed Asset"),
                ],
                help_text="The classification of the general ledger account.",
                max_length=19,
                verbose_name="Account Classification",
            ),
        ),
        migrations.AlterField(
            model_name="generalledgeraccount",
            name="account_sub_type",
            field=models.ChoiceField(
                blank=True,
                choices=[
                    ("CURRENT_ASSET", "Current Asset"),
                    ("FIXED_ASSET", "Fixed Asset"),
                    ("OTHER_ASSET", "Other Asset"),
                    ("CURRENT_LIABILITY", "Current Liability"),
                    ("LONG_TERM_LIABILITY", "Long Term Liability"),
                    ("EQUITY", "Equity"),
                    ("REVENUE", "Revenue"),
                    ("COST_OF_GOODS_SOLD", "Cost of Goods Sold"),
                    ("EXPENSE", "Expense"),
                    ("OTHER_INCOME", "Other Income"),
                    ("OTHER_EXPENSE", "Other Expense"),
                ],
                help_text="The sub type of the general ledger account.",
                max_length=19,
                verbose_name="Account Sub Type",
            ),
        ),
        migrations.AlterField(
            model_name="generalledgeraccount",
            name="account_type",
            field=models.ChoiceField(
                choices=[
                    ("ASSET", "Asset"),
                    ("LIABILITY", "Liability"),
                    ("EQUITY", "Equity"),
                    ("REVENUE", "Revenue"),
                    ("EXPENSE", "Expense"),
                ],
                help_text="The type of the general ledger account.",
                max_length=9,
                verbose_name="Account Type",
            ),
        ),
        migrations.AlterField(
            model_name="generalledgeraccount",
            name="cash_flow_type",
            field=models.ChoiceField(
                blank=True,
                choices=[
                    ("OPERATING", "Operating"),
                    ("INVESTING", "Investing"),
                    ("FINANCING", "Financing"),
                ],
                help_text="The cash flow type of the general ledger account.",
                max_length=9,
                verbose_name="Cash Flow Type",
            ),
        ),
    ]
