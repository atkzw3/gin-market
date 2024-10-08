package controllers

import (
	"gin-market/dto"
	"gin-market/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetById(ctx *gin.Context)
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

func (authController *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := authController.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed login"})
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (authController *AuthController) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := authController.service.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
