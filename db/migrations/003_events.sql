-- +goose Up
CREATE TABLE events (
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_code INT NOT NULL UNIQUE,
    PRIMARY KEY(id)
); 

-- +goose Down
DROP TABLE IF EXISTS events;
