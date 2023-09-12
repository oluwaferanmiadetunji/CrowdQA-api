-- +goose Up
CREATE TABLE polls (
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    poll_type VARCHAR(20) NOT NULL,
    poll_question TEXT,
    PRIMARY KEY(id)
); 

-- +goose Down
DROP TABLE IF EXISTS polls;
