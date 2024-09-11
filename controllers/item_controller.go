package controllers

import (
	"fmt"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IItemController interface {
	GetAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
type ItemController struct {
	service services.IItemService
}

func NewItemController(s services.IItemService) IItemController {
	return &ItemController{service: s}
}

func (ic *ItemController) GetAll(ctx *gin.Context) {
	items, err := ic.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (ic *ItemController) FindById(ctx *gin.Context) {
	// パスパラメーターはstring型なので、uint変換
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid id"})
	}
	item, err := ic.service.FindById(uint(id))
	fmt.Println("item", item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
