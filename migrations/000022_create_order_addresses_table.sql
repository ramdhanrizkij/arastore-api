-- +goose Up
CREATE TABLE IF NOT EXISTS order_addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    type address_type NOT NULL,
    recipient_name VARCHAR(150) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    address_line TEXT NOT NULL,
    province VARCHAR(150),
    city VARCHAR(150),
    district VARCHAR(150),
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL DEFAULT 'Indonesia',
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_order_addresses_order_id_type UNIQUE (order_id, type)
);

CREATE INDEX IF NOT EXISTS idx_order_addresses_order_id ON order_addresses(order_id);

-- +goose Down
DROP TABLE IF EXISTS order_addresses;
