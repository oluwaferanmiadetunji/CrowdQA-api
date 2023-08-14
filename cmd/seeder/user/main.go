package userSeeder

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/users"
)

var (
	logger, _ = thoth.Init("log")
)

func SeedUsers() {
	fmt.Println("Running user seeder...")

	for i := 0; i < 10; i++ {
		profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)

		name := randomdata.FullName(randomdata.RandomGender)

		params := users.UserParameters{
			Name:     name,
			Email:    randomdata.Email(),
			Password: profile.Login.Md5,
		}

		_, err := users.SaveUserToDb(params)

		if err != nil {
			logger.Log(fmt.Errorf("error creating user %s: %v", name, err))
			fmt.Printf("Error creating user %s: %v\n", name, err)
		} else {
			fmt.Printf("User %s created\n", name)
		}
	}
}
