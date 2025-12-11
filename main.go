package main

import (
	"Shop/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file not found, using environment variables instead")
	}

	db := config.Connect()
	config.Route(db)
}
