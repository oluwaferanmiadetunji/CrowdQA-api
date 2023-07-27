package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	EventCode int       `json:"event_code"`
}

func ConvertDatabaseEventToEvent(dbEvent database.Event) Event {
	return Event{
		ID:        dbEvent.ID,
		CreatedAt: dbEvent.CreatedAt,
		Name:      dbEvent.Name,
		StartDate: dbEvent.StartDate,
		EndDate:   dbEvent.EndDate,
		EventCode: int(dbEvent.EventCode),
	}
}

func ConvertDatabaseEventsToEvents(dbEvents []database.Event) []Event {
	events := []Event{}

	for _, dbEvent := range dbEvents {
		events = append(events, ConvertDatabaseEventToEvent(dbEvent))
	}

	return events
}
