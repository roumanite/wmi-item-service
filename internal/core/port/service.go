package port

import "wmi-item-service/internal/core/domain"

type UserService interface {
	CreateUser() error
}

type ResidenceService interface {
	CreateResidence(domain.CreateResidenceRequest) (*domain.Residence, error)
}