-- +goose Up
CREATE TABLE IF NOT EXISTS variant_attribute_values (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    variant_id UUID NOT NULL REFERENCES product_variants(id) ON DELETE CASCADE,
    attribute_value_id UUID NOT NULL REFERENCES attribute_values(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_variant_attribute_values UNIQUE (variant_id, attribute_value_id)
);

CREATE INDEX IF NOT EXISTS idx_variant_attribute_values_variant_id ON variant_attribute_values(variant_id);
CREATE INDEX IF NOT EXISTS idx_variant_attribute_values_attribute_value_id ON variant_attribute_values(attribute_value_id);

-- +goose Down
DROP TABLE IF EXISTS variant_attribute_values;
