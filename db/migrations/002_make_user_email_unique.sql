-- +goose Up
ALTER TABLE users ADD UNIQUE (email);


-- +goose Down
