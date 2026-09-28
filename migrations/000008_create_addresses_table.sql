-- +goose Up
CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    recipient_name VARCHAR(150) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    address_line TEXT NOT NULL,
    province VARCHAR(150),
    city VARCHAR(150),
    district VARCHAR(150),
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL DEFAULT 'Indonesia',
    notes TEXT,
    is_default BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_addresses_user_id ON addresses(user_id);

-- +goose Down
DROP TABLE IF EXISTS addresses;
