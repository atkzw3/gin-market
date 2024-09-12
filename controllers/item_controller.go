package controllers

import (
	"fmt"
	"gin-market/dto"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IItemController interface {
	GetAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
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

func (ic *ItemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput

	/*
		ShouldBindJSONについて
		https://qiita.com/ko-watanabe/items/64134c0a3871856fdc17
	*/
	//

	/*
		エラー
		invalid character '}' looking for beginning of object key string
		https://stackoverflow.com/questions/29690789/json-invalid-character-looking-for-beginning-of-object-key-string
	*/

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println("ShouldBindJSON エラー発生")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := ic.service.Create(input)
	if err != nil {
		fmt.Println("createエラー発生", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (ic *ItemController) Update(ctx *gin.Context) {
	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println("ShouldBindJSON エラー発生", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid id"})
	}

	fmt.Println("id:", uint(id))

	updatedItem, err := ic.service.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}
