package events

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/db"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

var (
	apiCfg    = db.Database()
	logger, _ = thoth.Init("log")
)

type EventParameters struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type PollParameters struct {
	PollQuestion string   `json:"poll_question"`
	PollType     string   `json:"poll_type"`
	EndDate      string   `json:"end_date"`
	PollOptions  []string `json:"poll_options"`
}

func CreateEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := EventParameters{}

	err := decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error creating event, please try again")
		return
	}

	event, err := SaveEventToDB(params, user.ID)

	if err != nil {
		logger.Log(fmt.Errorf("error creating event %v", err))
		log.Printf("error creating event: %v", err)
		utils.ErrorResponse(w, 400, ("Error creating event, please try again"))
		return
	}

	utils.JSONResponse(w, 201, utils.ConvertDatabaseEventToEvent(*event))

}

func GetMyEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	page, offset := utils.GetQueryOffset(r)

	events, err := apiCfg.DB.GetMyEvents(r.Context(), database.GetMyEventsParams{
		UserID: user.ID,
		Offset: offset,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events %v", err))
		log.Printf("error fetching events: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events, please try again"))
		return
	}

	eventCount, err := apiCfg.DB.GetMyEventsCount(r.Context(), user.ID)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events count %v", err))
		log.Printf("error fetching events count: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events count, please try again"))
		return
	}

	response := utils.APIQueryResponseStruct{
		Data:        events,
		Count:       eventCount,
		CurrentPage: page + 1,
		TotalPages:  utils.GetNumberOfPagesFromCount(eventCount),
		Limit:       10,
	}

	utils.JSONResponse(w, 200, response)
}

func GetMyUpcomingEvents(w http.ResponseWriter, r *http.Request, user database.User) {
	page, offset := utils.GetQueryOffset(r)

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

	eventCount, err := apiCfg.DB.GetUpComingEventsCount(r.Context(), user.ID)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events count %v", err))
		log.Printf("error fetching events count: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events count, please try again"))
		return
	}

	response := utils.APIQueryResponseStruct{
		Data:        events,
		Count:       eventCount,
		CurrentPage: page + 1,
		TotalPages:  utils.GetNumberOfPagesFromCount(eventCount),
		Limit:       10,
	}

	utils.JSONResponse(w, 200, response)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request, user database.User) {
	params := mux.Vars(r)
	eventId := params["id"]

	parsedEventId, err := uuid.Parse(eventId)

	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	err = apiCfg.DB.DeleteEvent(r.Context(), database.DeleteEventParams{
		ID:     parsedEventId,
		UserID: user.ID,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error deleting event %v", err))
		log.Printf("error deleting event: %v", err)
		utils.ErrorResponse(w, 400, ("Error deleting event, please try again"))
		return
	}

	response := utils.Response{
		Message: "Event deleted",
	}

	utils.JSONResponse(w, 200, response)
}

func GetMyEventById(w http.ResponseWriter, r *http.Request, user database.User) {
	params := mux.Vars(r)
	eventId := params["id"]

	parsedEventId, err := uuid.Parse(eventId)

	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	event, err := apiCfg.DB.GetEventById(r.Context(), database.GetEventByIdParams{
		ID:     parsedEventId,
		UserID: user.ID,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error fetching event %v", err))
		log.Printf("error fetching event: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching event, please try again"))
		return
	}

	utils.JSONResponse(w, 201, utils.ConvertDatabaseEventToEvent(event))
}

