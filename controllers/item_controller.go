package controllers

import (
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IItemController interface {
	GetAll(ctx *gin.Context)
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
