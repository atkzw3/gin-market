package controllers

import (
	"gin-market/dto"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (authController *AuthController) SignUp(ctx *gin.Context) {
	var input dto.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := authController.service.SignUp(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed sign up"})
	}
	ctx.Status(http.StatusCreated)
}
