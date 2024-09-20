package services

import (
	"fmt"
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
	GetByToken(token string) (*models.User, error)
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

func (a *AuthService) GetByToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 有効期限確認
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}
		user, err = a.r.FindByEmail(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
