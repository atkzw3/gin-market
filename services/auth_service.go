package services

import (
	"gin-market/models"
	"gin-market/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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

	token, err := CreateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	// jwtにおけるクレームは様々なユーザー情報を入れる
	// https://github.com/golang-jwt/jwt?tab=readme-ov-file#installation-guidelines
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	// envのsecret key は openssl rand -hex 32 でランダムな値を生成している
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// jwt 仕組みについて
	// https://jwt.io/

	return &tokenString, nil
}
