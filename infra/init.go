package infra

import (
	"github.com/joho/godotenv"
	"log"
)

func Initialize() {
	//go dot env package
	// https://github.com/joho/godotenv#installation
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
