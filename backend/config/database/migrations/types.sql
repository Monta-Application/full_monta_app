DO $$ 
BEGIN

-- Organization Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'org_type'
) THEN 
    CREATE TYPE org_type AS ENUM ('A', 'B', 'X');
END IF;

-- Generic Status Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'status_type'
) THEN 
    CREATE TYPE status_type AS ENUM ('A', 'I');
END IF;

-- Email Profile: Email Protocol Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'email_protocol_type'
) THEN 
    CREATE TYPE email_protocol_type AS ENUM ('TLS', 'SSL', 'UNENCRYPTED');
END IF;

-- Accounting Control: Automatic Journal Entry Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'automatic_journal_entry_type'
) THEN 
    CREATE TYPE automatic_journal_entry_type AS ENUM (
        'ON_SHIPMENT_BILL',
        'ON_RECEIPT_OF_PAYMENT',
        'ON_EXPENSE_RECOGNITION'
    );
END IF;

-- Accounting Control: Threshold Action Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'ac_threshold_action_type'
) THEN 
    CREATE TYPE ac_threshold_action_type AS ENUM (
        'HALT',
        'WARN'
    );
END IF;

-- General Ledger Account: Account Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'ac_account_type'
) THEN 
    CREATE TYPE ac_account_type AS ENUM (
        'ASSET',
        'LIABILITY',
        'EQUITY',
        'REVENUE',
        'EXPENSE'
    );
END IF;

-- General Ledger Account: Cash Flow Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'ac_cash_flow_type'
) THEN 
    CREATE TYPE ac_cash_flow_type AS ENUM (
        'OPERATING',
        'INVESTING',
        'FINANCING'
    );
END IF;

-- General Ledger Account: Account Sub Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'ac_account_sub_type'
) THEN 
    CREATE TYPE ac_account_sub_type AS ENUM (
        'CURRENT_ASSET',
        'FIXED_ASSET',
        'OTHER_ASSET',
        'CURRENT_LIABILITY',
        'LONG_TERM_LIABILITY',
        'EQUITY',
        'REVENUE',
        'COST_OF_GOODS_SOLD',
        'EXPENSE',
        'OTHER_INCOME',
        'OTHER_EXPENSE'
    );
END IF;

-- General Ledger Account: Account Classification --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'ac_account_classification'
) THEN 
    CREATE TYPE ac_account_classification AS ENUM (
        'BANK',
        'CASH',
        'ACCOUNTS_RECEIVABLE',
        'ACCOUNTS_PAYABLE',
        'INVENTORY',
        'OTHER_CURRENT_ASSET',
        'FIXED_ASSET'
    );
END IF;

-- Table Change Alert: Database Action Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'database_action_type'
) THEN 
    CREATE TYPE database_action_type AS ENUM ('INSERT', 'UPDATE', 'DELETE', 'ALL');
END IF;

-- Table Change Alert: Source Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'table_change_type'
) THEN 
    CREATE TYPE table_change_type AS ENUM ('KAFKA', 'DB');
END IF;

-- Job Title: Job Function Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'job_function_type'
) THEN 
    CREATE TYPE job_function_type AS ENUM (
        'MGR',
        'MT',
        'SP',
        'D',
        'B',
        'F',
        'S',
        'M',
        'A'
    );
END IF;

-- User: Timezone Type --
IF NOT EXISTS (
    SELECT
        1
    FROM
        pg_type
    WHERE
        typname = 'timezone_type'
) THEN 
    CREATE TYPE timezone_type AS ENUM (
        'America/Los_Angeles',
        'America/Denver',
        'America/Chicago',
        'America/New_York'
    );
END IF;

END $$;
