-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS profile_picture TEXT;

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS profile_picture;
