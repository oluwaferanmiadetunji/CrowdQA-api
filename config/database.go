package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	password = os.Getenv("DB_PASSWORD")
	user     = os.Getenv("DB_USER")
	port     = os.Getenv("DB_PORT")
	name     = os.Getenv("DB_NAME")
	host     = os.Getenv("DB_HOST")
)

func Database() *sql.DB {
	credentials := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True", user, password, host, port, name)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	// _, err = database.Exec(`CREATE DATABASE %s`, name)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = database.Exec(`USE %s`, name)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	return database
}
