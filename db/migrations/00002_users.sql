-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login text,
    password_hash text
);

-- +goose Down
DROP TABLE users;
