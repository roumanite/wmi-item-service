package service

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/core/port"
)

type ItemService struct {
	repo port.ItemRepository
}

func NewItemService(repo port.ItemRepository) *ItemService {
	return &ItemService{repo}
}

func (s *ItemService) CreateItem(req domain.CreateItemRequest) (*domain.Item, error) {
	item, err := s.repo.CreateItem(req)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// check *****
func (s *ItemService) UpdateItem(req domain.UpdateItemRequest) (*domain.Item, error) {
	item, err := s.repo.UpdateItem(req)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) GetItem(req domain.GetItemRequest) (*domain.Item, error) {
	item, err := s.repo.GetItem(req)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) DeleteItem(req domain.DeleteItemRequest) (error) {
	err := s.repo.DeleteItem(req)
	if err != nil {
		return err
	}
	return nil
}