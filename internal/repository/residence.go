package repository

import (
	"fmt"
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
)

type ResidenceRepo struct {
	db *gorm.DB
}

func NewResidenceRepo(db *gorm.DB) *ResidenceRepo {
	return &ResidenceRepo{db}
}

func (r *ResidenceRepo) CreateResidence(req domain.CreateResidenceRequest) (*domain.Residence, error) {
	residence := domain.Residence{
		UserId: req.UserId,
		Nickname: req.Nickname,
		StreetAddress: req.StreetAddress,
		City: req.City,
		State: req.State,
		Country: req.Country,
		ZipCode: req.ZipCode,
		BuildingName: req.BuildingName,
	}
	err := r.db.Create(&residence).Error
	if err != nil {
		fmt.Printf("create residence db error %v\n", err)
		return nil, err
	}
	return &residence, err
}