-- +goose Up
ALTER TABLE poll_options
ALTER COLUMN option_text SET NOT NULL;
