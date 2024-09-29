package repositories

import (
	"gin-market/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindByEmail(email string) (*models.User, error)
	GetById(id uint) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
