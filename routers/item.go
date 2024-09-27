package routers

import (
	"gin-market/controllers"
	"gin-market/middlewares"
	"gin-market/repositories"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupItemRoutes(r *gin.Engine, db *gorm.DB, service services.IAuthService) *gin.Engine {
	ic := getItemController(db)
	itemR := r.Group("/items")
	itemWithAuth := r.Group("/items", middlewares.AuthMiddleware(service))
	itemR.GET("", ic.GetAll)
	// TODO デバックのため、一旦middleware外す
	itemR.GET("/:id", ic.FindById)
	itemWithAuth.POST("", ic.Create)
	itemWithAuth.PUT("/:id", ic.Update)
	itemWithAuth.DELETE("/:id", ic.Delete)

	return r
}

func getItemController(db *gorm.DB) controllers.IItemController {
	ir := repositories.NewDBItemRepository(db)
	is := services.NewItemService(ir)
	ic := controllers.NewItemController(is)

	return ic
}
