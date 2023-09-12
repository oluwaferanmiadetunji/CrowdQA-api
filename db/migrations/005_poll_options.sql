-- +goose Up
CREATE TABLE poll_options (
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    poll_id UUID NOT NULL REFERENCES polls(id) ON DELETE CASCADE,
    option_text TEXT,
    PRIMARY KEY(id)
); 

-- +goose Down
DROP TABLE IF EXISTS poll_options;
