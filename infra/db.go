package infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func SetupDB() *gorm.DB {

	env := os.Getenv("ENV")

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	if env == "prod" || env == "local" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("set up postgres connection")
		if err != nil {
			panic("failed db connect")
		}
		return db
	} else {
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		log.Println("set up sqlite connection")
		if err != nil {
			panic("failed db connect")
		}
		return db
	}
}
