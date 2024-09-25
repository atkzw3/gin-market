package main

import (
	"bytes"
	"encoding/json"
	"gin-market/dto"
	"gin-market/infra"
	"gin-market/models"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestGetAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	// APIリクエスト実行
	router.ServeHTTP(w, req)

	// 実行結果取得
	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)

	// body
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@test.com")
	// error がないこと確認
	assert.Nil(t, err)

	createItemInput := dto.CreateItemInput{
		Name:        "test item 4",
		Price:       4000,
		Description: "test description 4",
	}

	// json エンコード
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// APIリクエスト実行
	router.ServeHTTP(w, req)

	// 実行結果取得
	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)

	// body
	assert.Equal(t, uint(4), res["data"].ID)
}
