package eventsSeeder

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/events"
)

var (
	logger, _ = thoth.Init("log")
)

func generateRandomName() string {
	names := []string{"Conference", "Seminar", "Workshop", "Webinar", "Symposium", "Meetup"}
	return names[rand.Intn(len(names))]
}

func generateRandomDate() string {
	min := time.Now()
	max := min.AddDate(0, 1, 0) // 1 month ahead
	randomTime := min.Add(time.Duration(rand.Int63n(max.Unix()-min.Unix())) * time.Second)
	return randomTime.Format(time.RFC3339)

}

func generateRandomDateAfter(after string) string {
	afterTime, _ := time.Parse(time.RFC3339, after)
	min := afterTime
	max := min.AddDate(0, 0, 1) // 1 day ahead
	randomTime := min.Add(time.Duration(rand.Int63n(max.Unix()-min.Unix())) * time.Second)
	return randomTime.Format(time.RFC3339)
}

func SeedEvents(userID uuid.UUID) {
	fmt.Printf("Running event seeder for user ID: %s \n\n", userID)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		randomEventType := generateRandomName()
		randomEventName := randomdata.SillyName()

		name := randomEventName + " " + randomEventType

		startDate := generateRandomDate()
		endDate := generateRandomDateAfter(startDate)

		params := events.EventParameters{
			Name:      name,
			StartDate: startDate,
			EndDate:   endDate,
		}

		_, err := events.SaveEventToDB(params, userID)

		if err != nil {
			logger.Log(fmt.Errorf("error creating user %s: %v", name, err))
			fmt.Printf("Error creating user %s: %v\n", name, err)
		} else {
			fmt.Printf("Event %s created\n", name)
		}
	}
}
