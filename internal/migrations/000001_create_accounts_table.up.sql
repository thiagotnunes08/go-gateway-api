CREATE TABLE IF NOT EXISTS accounts (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL UNIQUE,
api_key VARCHAR(255) NOT NULL UNIQUE,
balance decimal(10,2) NOT NULL DEFAULT 0,
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_accounts_api_key on accounts(api_key);

CREATE INDEX idx_accounts_email on accounts(email);

CREATE TABLE IF NOT EXISTS invoices (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
account_id UUID NOT NULL REFERENCES accounts(id),
amount decimal(10,2) NOT NULL DEFAULT 0,
status VARCHAR(50) NOT NULL DEFAULT 'pending',
description TEXT NOT NULL,
payment_type VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_invoices_account_id on invoices(account_id);

CREATE INDEX idx_invoices_status on invoices(status);

CREATE INDEX idx_invoices_created_at on invoices(created_at);