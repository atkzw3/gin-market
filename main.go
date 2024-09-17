package main

import (
	"gin-market/controllers"
	"gin-market/infra"
	"gin-market/repositories"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// https://gin-gonic.com/ja/docs/quickstart/

// air パッケージでホットリロード
// https://github.com/air-verse/air
func main() {
	infra.Initialize()
	db := infra.SetupDB()
	log.Println(os.Getenv("ENV"))
	//
	//items := []models.Item{
	//	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	//	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	//	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	//}
	//ir := repositories.NewItemRepository(items)

	ir := repositories.NewDBItemRepository(db)
	is := services.NewItemService(ir)
	ic := controllers.NewItemController(is)

	r := gin.Default()
	itemR := r.Group("/items")

	itemR.GET("/", ic.GetAll)
	itemR.GET("/:id", ic.FindById)
	itemR.POST("/", ic.Create)
	itemR.PUT("/:id", ic.Update)
	itemR.DELETE("/:id", ic.Delete)

	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
