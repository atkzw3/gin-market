package main

import (
	"gin-market/infra"
	"gin-market/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	code := m.Run()
	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "test item1", Price: 1000, Description: "test description 1", SoldOut: false, UserID: 1},
		{Name: "test item2", Price: 2000, Description: "test description 2", SoldOut: true, UserID: 1},
		{Name: "test item3", Price: 3000, Description: "test description 3", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@test.com", Password: "test1pass"},
		{Email: "test2@test.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)

	router := setupRouter(db)
	return router
}
