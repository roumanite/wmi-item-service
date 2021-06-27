package port

import "wmi-item-service/internal/core/domain"

type UserRepository interface {
	CreateUser(domain.CreateUserRequest) error
	GetUser(identifier string) (*domain.User, error)
	GetProfile(id string) (*domain.UserProfile, error)
	UpdateProfile(domain.UpdateProfileRequest) (*domain.UserProfile, error)
}

type ResidenceRepository interface {
	CreateResidence(domain.CreateResidenceRequest) (*domain.Residence, error)
	UpdateResidence(domain.UpdateResidenceRequest) (*domain.Residence, error)
	GetResidence(domain.GetResidenceRequest) (*domain.Residence, error)
	GetResidenceList(domain.GetResidenceListRequest) (*domain.MetaResidences, error)
	DeleteResidence(domain.DeleteResidenceRequest) (error)
}

type ItemRepository interface {
	CreateItem(domain.CreateItemRequest) (*domain.Item, error)
	UpdateItem(domain.UpdateItemRequest) (*domain.Item, error)
	GetItem(domain.GetItemRequest) (*domain.Item, error)
	GetItemList(domain.GetItemListRequest) (*domain.MetaItems, error)
	DeleteItem(domain.DeleteItemRequest) (error)
}