package polls

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	database "github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

func SavePollToDB(payload database.CreatePollParams) (*database.Poll, error) {
	ctx := context.Background()

	poll, err := apiCfg.DB.CreatePoll(ctx, database.CreatePollParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		EventID:      payload.EventID,
		PollType:     payload.PollType,
		PollQuestion: payload.PollQuestion,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating poll %v", err))
		log.Printf("error creating poll: %v", err)

		return nil, err
	}

	return &poll, nil
}

func SavePollOptionsToDB(payload database.CreatePollOptionsParams) (*database.PollOption, error) {
	ctx := context.Background()

	pollOptions, err := apiCfg.DB.CreatePollOptions(ctx, database.CreatePollOptionsParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		PollID:     payload.PollID,
		OptionText: payload.OptionText,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating poll options %v", err))
		log.Printf("error creating poll options: %v", err)

		return nil, err
	}

	return &pollOptions, nil
}

func GetPollsByEventId(event_id uuid.UUID) (*[]database.Poll, error) {
	ctx := context.Background()

	polls, err := apiCfg.DB.GetPollsByEventId(ctx, event_id)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching polls %v", err))
		log.Printf("error fetching polls: %v", err)

		return nil, err
	}

	return &polls, nil
}
