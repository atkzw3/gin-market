package repositories

import "gin-market/models"

type IItemRepository interface {
	GetAll() (*[]models.Item, error)
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
