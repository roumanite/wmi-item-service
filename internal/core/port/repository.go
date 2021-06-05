package port

import "wmi-item-service/internal/core/domain"

type UserRepository interface {
	CreateUser() error
}

type ResidenceRepository interface {
	CreateResidence(domain.CreateResidenceRequest) (*domain.Residence, error)
}

type ItemRepository interface {
	CreateItem(domain.CreateItemRequest) (*domain.Item, error)
}