package events

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

func SaveEventToDB(params EventParameters, user_id uuid.UUID) (*database.Event, error) {
	ctx := context.Background()
	event_code, err := utils.GenerateEventCode()

	if err != nil {
		logger.Log(fmt.Errorf("error generating event code %v", err))
		log.Printf("error generating event code: %v", err)
		return nil, err
	}

	event, err := apiCfg.DB.CreateEvent(ctx, database.CreateEventParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		StartDate: utils.ConvertStringToTime(params.StartDate),
		EndDate:   utils.ConvertStringToTime(params.EndDate),
		UserID:    user_id,
		EventCode: event_code,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating event %v", err))
		log.Printf("error creating event: %v", err)

		return nil, err
	}

	return &event, nil
}
