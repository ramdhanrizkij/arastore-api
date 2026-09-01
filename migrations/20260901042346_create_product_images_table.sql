-- +goose Up
CREATE TABLE IF NOT EXISTS product_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES products(id),
    media_id UUID NOT NULL REFERENCES media(id),
    is_primary BOOLEAN DEFAULT FALSE,
    sort_order int default 0,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_product_images_id
ON product_images(product_id);

-- +goose Down
DROP TABLE IF EXISTS product_images;
