package polls

import (
	"context"
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

type PollParameters struct {
	PollQuestion string   `json:"poll_question"`
	PollType     string   `json:"poll_type"`
	PollOptions  []string `json:"poll_options"`
}

func CreatePoll(w http.ResponseWriter, r *http.Request, user database.User) {
	routeParams := mux.Vars(r)
	eventId := routeParams["id"]

	parsedEventId, err := uuid.Parse(eventId)

	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := PollParameters{}

	err = decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error creating event, please try again")
		return
	}

	poll, err := SavePollToDB(database.CreatePollParams{
		EventID:      parsedEventId,
		PollType:     params.PollType,
		PollQuestion: params.PollQuestion,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating poll %v", err))
		log.Printf("error creating poll: %v", err)
		utils.ErrorResponse(w, 400, ("Error creating poll, please try again"))
		return
	}

	var savedOptions []*database.PollOption

	for i := 0; i < len(params.PollOptions); i++ {

		option, err := SavePollOptionsToDB(database.CreatePollOptionsParams{
			PollID:     poll.ID,
			OptionText: params.PollOptions[i],
		})

		if err != nil {
			logger.Log(fmt.Errorf("error creating poll option %v", err))
			log.Printf("error creating poll option: %v", err)
		} else {
			savedOptions = append(savedOptions, option)
		}
	}

	utils.JSONResponse(w, 201, utils.AddPollOptionsToPoll(*poll, savedOptions))
}

func FetchPolls(w http.ResponseWriter, r *http.Request, user database.User) {
	ctx := context.Background()
	params := mux.Vars(r)
	eventId := params["id"]

	parsedEventId, err := uuid.Parse(eventId)

	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	polls, err := GetPollsByEventId(parsedEventId)

	if err != nil {
		logger.Log(fmt.Errorf("error fetching polls %v", err))
		log.Printf("error fetching polls: %v", err)
		utils.ErrorResponse(w, 400, ("Error fetching polls, please try again"))
		return
	}

	allPolls := []utils.Poll{}

	for _, poll := range *polls {

		option, err := apiCfg.DB.GetPollOptionsByPollId(ctx, poll.ID)

		if err != nil {
			logger.Log(fmt.Errorf("error creating poll option %v", err))
			log.Printf("error creating poll option: %v", err)
		} else {
			var optionPointers []*database.PollOption
			for i := range option {
				optionPointers = append(optionPointers, &option[i])
			}

			newPoll := utils.AddPollOptionsToPoll(poll, optionPointers)

			allPolls = append(allPolls, newPoll)
		}
	}

	utils.JSONResponse(w, 200, allPolls)
}
