package repositories

import (
	"errors"
	"gin-market/models"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(deleteItem models.Item) error
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

func (r *ItemRepositoryImpl) Update(updateItem models.Item) (*models.Item, error) {
	for i, item := range r.items {
		if item.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemRepositoryImpl) Delete(deleteItem models.Item) error {
	for i, item := range r.items {
		if item.ID == deleteItem.ID {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func (i *ItemRepository) GetAll() (*[]models.Item, error) {
	var items []models.Item
	r := i.db.Find(&items)
	if r.Error != nil {
		return nil, r.Error
	}
	return &items, r.Error
}

func (i *ItemRepository) FindById(id uint) (*models.Item, error) {
	var item models.Item
	r := i.db.First(&item, id)
	if r.Error != nil {
		return nil, r.Error
	}
	return &item, nil
}

func (i *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	r := i.db.Create(&newItem)
	// gormの場合、返り値にerrorが含まれる
	if r.Error != nil {
		return nil, r.Error
	}
	return &newItem, nil
}

func (i *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	r := i.db.Save(&updateItem)
	if r.Error != nil {
		return nil, r.Error
	}
	return &updateItem, nil
}

func (i ItemRepository) Delete(deleteItem models.Item) error {
	//TODO implement me
	panic("implement me")
}

func NewDBItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}
