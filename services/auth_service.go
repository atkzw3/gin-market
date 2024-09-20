package services

import (
	"gin-market/models"
	"gin-market/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
	Login(email string, password string) (*string, error)
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

func (a *AuthService) Login(email string, password string) (*string, error) {
	user, err := a.r.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user.Email, nil
}
