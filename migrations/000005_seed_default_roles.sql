-- +goose Up
INSERT INTO roles (name, description) VALUES
    ('superadmin', 'Super Administrator with full access'),
    ('admin', 'Administrator'),
    ('user', 'Regular user')
ON CONFLICT (name) DO NOTHING;

-- +goose Down
DELETE FROM roles WHERE name IN ('superadmin', 'admin', 'user');
