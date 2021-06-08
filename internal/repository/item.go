package repository

import (
	"fmt"
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
	"time"
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

// check *****
func (r *ItemRepo) UpdateItem(req domain.UpdateItemRequest) (*domain.Item, error) {
	item := domain.Item{}
	err := r.db.Model(&item).Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Updates(map[string]interface{}{
		"name": req.Name,
		"category_id": req.CategoryId,
		"display_picture_url": req.DisplayPictureUrl,
		"notes": req.Notes,
	}).Take(&item).Error // TODO
	if err != nil {
		fmt.Printf("update item db error %v\n", err)
		return nil, err
	}
	return &item, err
}

func (r *ItemRepo) GetItem(req domain.GetItemRequest) (*domain.Item, error) {
	item := domain.Item{}
	err := r.db.Table("items").Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Take(&item).Error // TODO
	if err != nil {
		fmt.Printf("get item db error %v\n", err)
		return nil, err
	}
	return &item, err
}

func (r *ItemRepo) DeleteItem(req domain.DeleteItemRequest) (error) {
	err := r.db.Table("items").Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error // TODO
	if err != nil {
		fmt.Printf("delete item db error %v\n", err)
		return err
	}
	return nil
}