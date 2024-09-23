package main

import (
	"gin-market/controllers"
	"gin-market/infra"
	"gin-market/middlewares"
	"gin-market/repositories"
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
	ir := repositories.NewDBItemRepository(db)
	is := services.NewItemService(ir)
	ic := controllers.NewItemController(is)

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	// https://github.com/gin-contrib/cors
	// 一旦デフォルトで全て許可。案件に応じて変更
	r.Use(cors.Default())
	itemR := r.Group("/items")
	itemWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authR := r.Group("/auth")

	itemR.GET("/", ic.GetAll)
	itemWithAuth.GET("/:id", ic.FindById)
	itemWithAuth.POST("", ic.Create)
	itemWithAuth.PUT("/:id", ic.Update)
	itemWithAuth.DELETE("/:id", ic.Delete)

	authR.POST("/signup", authController.SignUp)
	authR.POST("/login", authController.Login)

	return r
}
