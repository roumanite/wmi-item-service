package repository

import (
	"fmt"
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
)

type ItemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) *ItemRepo {
	return &ItemRepo{db}
}

func (r *ItemRepo) CreateItem(req domain.CreateItemRequest) (*domain.Item, error) {
	item := domain.Item{
		Name: req.Name,
		UserIdOwner: req.UserIdOwner,
		CategoryId: req.CategoryId,
		DisplayPictureUrl: req.DisplayPictureUrl,
		Notes: req.Notes,
	}
	err := r.db.Create(&item).Error
	if err != nil {
		fmt.Printf("create item db error %v\n", err)
		return nil, err
	}
	return &item, err
}