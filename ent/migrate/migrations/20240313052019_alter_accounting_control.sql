-- Modify "accounting_controls" table
ALTER TABLE "accounting_controls" ADD CONSTRAINT "accounting_controls_general_ledger_accounts_default_exp_account" FOREIGN KEY ("default_exp_account_id") REFERENCES "general_ledger_accounts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "accounting_controls_general_ledger_accounts_default_rev_account" FOREIGN KEY ("default_rev_account_id") REFERENCES "general_ledger_accounts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
