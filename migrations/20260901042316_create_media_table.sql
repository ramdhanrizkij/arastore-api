-- +goose Up
CREATE TABLE IF NOT EXISTS media(
    id UUID PRIMARY KEY,
    bucket_name VARCHAR(100) NOT NULL,
    object_key VARCHAR(500) NOT NULL,
    file_name VARCHAR(255),
    mime_type VARCHAR(100),
    file_size BIGINT,
    width int,
    height int,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS media;
