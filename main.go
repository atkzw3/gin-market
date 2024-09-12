package main

import (
	"gin-market/controllers"
	"gin-market/models"
	"gin-market/repositories"
	"gin-market/services"
	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/ja/docs/quickstart/

// air パッケージでホットリロード
// https://github.com/air-verse/air
func main() {

	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	ir := repositories.NewItemRepository(items)
	is := services.NewItemService(ir)
	ic := controllers.NewItemController(is)

	r := gin.Default()
	r.GET("/items", ic.GetAll)
	r.GET("/items/:id", ic.FindById)
	r.POST("/items", ic.Create)
	r.PUT("/items/:id", ic.Update)
	r.DELETE("/items/:id", ic.Delete)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
