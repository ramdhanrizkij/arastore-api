-- +goose Up
CREATE TABLE IF NOT EXISTS payment_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id UUID NOT NULL REFERENCES payments(id) ON DELETE CASCADE,
    transaction_id VARCHAR(255),
    reference_id VARCHAR(255),
    status payment_transaction_status NOT NULL DEFAULT 'pending',
    amount NUMERIC(19,2) NOT NULL CHECK (amount >= 0),
    payment_method VARCHAR(100),
    response_code VARCHAR(100),
    response_message TEXT,
    raw_request JSONB,
    raw_response JSONB,
    processed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_payment_transactions_payment_id ON payment_transactions(payment_id);
CREATE INDEX IF NOT EXISTS idx_payment_transactions_transaction_id ON payment_transactions(transaction_id);
CREATE INDEX IF NOT EXISTS idx_payment_transactions_reference_id ON payment_transactions(reference_id);
CREATE INDEX IF NOT EXISTS idx_payment_transactions_status ON payment_transactions(status);

-- +goose Down
DROP TABLE IF EXISTS payment_transactions;
