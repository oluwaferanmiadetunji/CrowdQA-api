package events

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/db"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

var (
	apiCfg    = db.Database()
	logger, _ = thoth.Init("log")
)

func CreateEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error creating event, please try again")
		return
	}

	event_code, err := utils.GenerateEventCode()

	if err != nil {
		logger.Log(fmt.Errorf("error generating event code %v", err))
		log.Printf("error generating event code: %v", err)
		utils.ErrorResponse(w, 400, "Error creating event, please try again")
		return
	}

	event, err := apiCfg.DB.CreateEvent(r.Context(), database.CreateEventParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		StartDate: utils.ConvertStringToTime(params.StartDate),
		EndDate:   utils.ConvertStringToTime(params.EndDate),
		UserID:    user.ID,
		EventCode: event_code,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating event %v", err))
		log.Printf("error creating event: %v", err)
		utils.ErrorResponse(w, 400, ("Error creating event, please try again"))
		return
	}

	utils.JSONResponse(w, 201, utils.ConvertDatabaseEventToEvent(event))

}

func GetMyEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	events, err := apiCfg.DB.GetMyEvents(r.Context(), user.ID)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events %v", err))
		log.Printf("error fetching events: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events, please try again"))
		return
	}

	utils.JSONResponse(w, 200, utils.ConvertDatabaseEventsToEvents(events))
}

func GetMyUpcomingEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	offsetStr := r.URL.Query().Get("offset")

	var offset int32

	if offsetStr != "" {
		parsedOffset, err := strconv.ParseInt(offsetStr, 10, 32)
		if err != nil {
			offset = 0
		} else {
			offset = int32(parsedOffset)
		}
	} else {
		offset = 0
	}

	events, err := apiCfg.DB.GetUpcomingEvents(r.Context(), database.GetUpcomingEventsParams{
		UserID: user.ID,
		Offset: offset,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events %v", err))
		log.Printf("error fetching events: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events, please try again"))
		return
	}

	utils.JSONResponse(w, 200, utils.ConvertDatabaseEventsToEvents(events))
}

func SeedEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	rand.Seed(time.Now().UnixNano())

	numEvents := 10

	for i := 0; i < numEvents; i++ {
		startDate := time.Now().Add(time.Duration(rand.Intn(7)) * 24 * time.Hour) // Random start date within the next 7 days
		endDate := startDate.Add(24 * time.Hour)

		event_code, err := utils.GenerateEventCode()

		if err != nil {
			logger.Log(fmt.Errorf("error generating event code %v", err))
			log.Printf("error generating event code: %v", err)
			utils.ErrorResponse(w, 400, "Error creating event, please try again")
			return
		}

		event, err := apiCfg.DB.CreateEvent(r.Context(), database.CreateEventParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      "params.Name",
			StartDate: startDate,
			EndDate:   endDate,
			UserID:    user.ID,
			EventCode: event_code,
		})

		if err != nil {
			logger.Log(fmt.Errorf("error creating event %v", err))
			log.Printf("error creating event: %v", err)
			utils.ErrorResponse(w, 400, ("Error creating event, please try again"))
			return
		}

		fmt.Printf(event.Name)
	}

	utils.JSONResponse(w, 201, "Event seeded successfully")
}
