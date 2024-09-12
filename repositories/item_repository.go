package repositories

import (
	"errors"
	"gin-market/models"
)

type IItemRepository interface {
	GetAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
}

type ItemRepositoryImpl struct {
	items []models.Item
}

func NewItemRepository(items []models.Item) IItemRepository {
	return &ItemRepositoryImpl{items: items}
}

func (r *ItemRepositoryImpl) GetAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemRepositoryImpl) FindById(id uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemRepositoryImpl) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}
