package services

import (
	"gin-market/models"
	"gin-market/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
}

type AuthService struct {
	r repositories.IAuthRepository
}

func NewAuthService(r repositories.IAuthRepository) IAuthService {
	return &AuthService{r: r}
}

func (a *AuthService) SignUp(email string, password string) error {
	hashP, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashP),
	}

	return a.r.CreateUser(user)
}
