package services

import (
	"fmt"
	"gin-market/dto"
	"gin-market/models"
	"gin-market/repositories"
)

type IItemService interface {
	GetAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(input dto.CreateItemInput, userId uint) (*models.Item, error)
	Update(id uint, input dto.UpdateItemInput) (*models.Item, error)
	Delete(id uint) error
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

func (s *ItemService) Create(input dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		SoldOut:     false,
		UserID:      userId,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(id uint, input dto.UpdateItemInput) (*models.Item, error) {
	item, err := s.FindById(id)
	if err != nil {
		fmt.Println("Update 時に該当データなし")
		return nil, err
	}

	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.Price != nil {
		item.Price = *input.Price
	}
	if input.Description != nil {
		item.Description = *input.Description
	}
	if input.SoldOut != nil {
		item.SoldOut = *input.SoldOut
	}

	return s.repository.Update(*item)
}

func (s *ItemService) Delete(id uint) error {
	item, err := s.FindById(id)
	if err != nil {
		fmt.Println("delete 時に該当データなし")
		return err
	}

	err2 := s.repository.Delete(*item)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	return nil
}
