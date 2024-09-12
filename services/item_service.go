package services

import (
	"gin-market/dto"
	"gin-market/models"
	"gin-market/repositories"
)

type IItemService interface {
	GetAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(input dto.CreateItemInput) (*models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) GetAll() (*[]models.Item, error) {
	return s.repository.GetAll()
}

func (s *ItemService) FindById(id uint) (*models.Item, error) {
	return s.repository.FindById(id)
}

func (s *ItemService) Create(input dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}
