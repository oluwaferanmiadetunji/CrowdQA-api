package events

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	pageStr := r.URL.Query().Get("page")

	var page int32
	var offset int32

	if pageStr != "" {
		parsedPage, err := strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			page = 0
		} else {
			page = int32(parsedPage)
		}
	} else {
		page = 0
	}

	if page == 0 {
		offset = 0
	} else {
		offset = page * 10
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

	eventCount, err := apiCfg.DB.GetUpComingEventsCount(r.Context(), user.ID)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching events count %v", err))
		log.Printf("error fetching events count: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching events count, please try again"))
		return
	}

	response := utils.APIQueryResponseStruct{
		Data:         events,
		Count:        eventCount,
		CurrentPage:  page + 1,
		TotalPages:   utils.GetNumberOfPagesFromCount(eventCount),
		Limit:        10,
	}

	utils.JSONResponse(w, 200, response)
}
