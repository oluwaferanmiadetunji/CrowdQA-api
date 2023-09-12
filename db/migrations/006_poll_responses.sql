-- +goose Up
CREATE TABLE poll_responses (
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    poll_id UUID NOT NULL REFERENCES polls(id) ON DELETE CASCADE,
    user_id TEXT,
    selected_option INT,
    user_answer TEXT,
    user_ranking INT[],
    user_rating INT,
    PRIMARY KEY(id)
); 

-- +goose Down
DROP TABLE IF EXISTS poll_responses;
