package main

import (
	"gin-market/controllers"
	"gin-market/infra"
	"gin-market/repositories"
	"gin-market/routers"
	"gin-market/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// https://gin-gonic.com/ja/docs/quickstart/

// air パッケージでホットリロード
// https://github.com/air-verse/air
func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := setupRouter(db)

	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	routers.SetupItemRoutes(r, db, authService)
	// https://github.com/gin-contrib/cors
	// 一旦デフォルトで全て許可。案件に応じて変更
	r.Use(cors.Default())
	authR := r.Group("/auth")

	authR.POST("/signup", authController.SignUp)
	authR.POST("/login", authController.Login)

	return r
}
