package repository

import (
	"wmi-item-service/pkg/postgresql"
	"wmi-item-service/internal/core/domain"
)

type ResidenceRepo struct {
	db *postgresql.Connection
}

func NewResidenceRepo(db *postgresql.Connection) *ResidenceRepo {
	return &ResidenceRepo{db}
}

func (r *ResidenceRepo) CreateResidence(req domain.CreateResidenceRequest) error {
	// TODO
	return nil
}