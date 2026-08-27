-- +goose Up

-- +goose StatementBegin
DO $$
BEGIN
    CREATE TYPE product_status AS ENUM (
        'DRAFT',
        'ACTIVE',
        'INACTIVE',
        'ARCHIVED'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id UUID NOT NULL REFERENCES categories(id),
    sku VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(100),
    description TEXT,
    price NUMERIC(18,2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    weight NUMERIC(10,2),
    status product_status NOT NULL DEFAULT 'DRAFT',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_products_category_id
ON products(category_id);

-- +goose Down

DROP TABLE IF EXISTS products;
DROP TYPE IF EXISTS product_status;