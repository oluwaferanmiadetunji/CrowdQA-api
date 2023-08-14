package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	eventsSeeder "github.com/oluwaferanmiadetunji/CrowdQA-api/cmd/seeder/events"
	userSeeder "github.com/oluwaferanmiadetunji/CrowdQA-api/cmd/seeder/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	seedCommand := flag.String("seed", "", "Specify the seeder to run (user or event)")
	userID := flag.String("user_id", "", "Specify the user ID")
	flag.Parse()

	if *seedCommand == "" {
		fmt.Println("Usage: seeder -seed=<user|event>")
		os.Exit(1)
	}

	switch *seedCommand {
	case "user":
		userSeeder.SeedUsers()

	case "event":
		eventsSeeder.SeedEvents(*userID)

	default:
		fmt.Printf("Invalid seeder: %s\n", *seedCommand)
		os.Exit(1)
	}

	fmt.Println("Seeder complete!")

}
