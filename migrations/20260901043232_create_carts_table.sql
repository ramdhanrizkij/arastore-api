-- +goose Up

-- +goose StatementBegin
DO $$
BEGIN
    CREATE TYPE cart_status AS ENUM (
        'ACTIVE',
        'CONVERTED',
        'ABANDONED'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS carts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    session_token varchar(256),
    status cart_status DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS carts;
DROP TYPE IF EXISTS cart_status;
