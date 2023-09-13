package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

type Poll struct {
	ID           uuid.UUID    `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	PollType     string       `json:"type"`
	PollQuestion string       `json:"question"`
	PollOptions  []PollOption `json:"options"`
}

type PollOption struct {
	ID         uuid.UUID `json:"id"`
	OptionText string    `json:"option_text"`
}

func AddPollOptionsToPoll(dbPoll database.Poll, pollOptions []*database.PollOption) Poll {
	return Poll{
		ID:           dbPoll.ID,
		CreatedAt:    dbPoll.CreatedAt,
		PollType:     dbPoll.PollType,
		PollQuestion: dbPoll.PollQuestion,
		PollOptions:  FormatPollOptions(pollOptions),
	}
}

func FormatPollOptions(dbPollOptions []*database.PollOption) []PollOption {
	options := []PollOption{}

	for _, option := range dbPollOptions {
		options = append(options, ConvertDBPollOptionToPollOption(*option))
	}

	return options
}

func ConvertDBPollOptionToPollOption(pollOption database.PollOption) PollOption {
	return PollOption{
		ID:         pollOption.ID,
		OptionText: pollOption.OptionText,
	}
}
