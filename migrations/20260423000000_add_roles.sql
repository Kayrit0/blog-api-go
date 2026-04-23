-- +goose Up
CREATE TYPE user_role AS ENUM ('user', 'admin', 'owner');

ALTER TABLE users
ADD COLUMN role user_role NOT NULL DEFAULT 'user';

CREATE INDEX idx_users_role ON users(role);

-- +goose Down
DROP INDEX IF EXISTS idx_users_role;
ALTER TABLE users DROP COLUMN role;
DROP TYPE user_role;
