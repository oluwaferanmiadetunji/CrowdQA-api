-- +goose Up
ALTER TABLE polls
ALTER COLUMN poll_question SET NOT NULL;
