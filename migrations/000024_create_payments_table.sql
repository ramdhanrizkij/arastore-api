-- +goose Up
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    gateway payment_gateway NOT NULL,
    status payment_status NOT NULL DEFAULT 'pending',
    amount NUMERIC(19,2) NOT NULL CHECK (amount >= 0),
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    external_payment_id VARCHAR(255),
    external_order_id VARCHAR(255),
    payment_method VARCHAR(100),
    expired_at TIMESTAMPTZ,
    paid_at TIMESTAMPTZ,
    raw_response JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_payments_order_id ON payments(order_id);
CREATE INDEX IF NOT EXISTS idx_payments_gateway ON payments(gateway);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);
CREATE INDEX IF NOT EXISTS idx_payments_external_payment_id ON payments(external_payment_id);
CREATE INDEX IF NOT EXISTS idx_payments_external_order_id ON payments(external_order_id);

-- +goose Down
DROP TABLE IF EXISTS payments;
