package repository

import (
	"fmt"
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
	"time"
)

type ResidenceRepo struct {
	db *gorm.DB
}

func NewResidenceRepo(db *gorm.DB) *ResidenceRepo {
	return &ResidenceRepo{db}
}

func (r *ResidenceRepo) CreateResidence(req domain.CreateResidenceRequest) (*domain.Residence, error) {
	residence := domain.Residence{
		UserIdOwner: req.UserIdOwner,
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

// check *****
func (r *ResidenceRepo) UpdateResidence(req domain.UpdateResidenceRequest) (*domain.Residence, error) {
	residence := domain.Residence{}
	err := r.db.Model(&residence).Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Updates(map[string]interface{}{
		"nickname": req.Nickname,
		"street_address": req.StreetAddress,
		"city": req.City,
		"state": req.State,
		"country": req.Country,
		"zip_code": req.ZipCode,
		"building_name": req.BuildingName,
	}).Take(&residence).Error // TODO
	if err != nil {
		fmt.Printf("update residence db error %v\n", err)
		return nil, err
	}
	return &residence, err
}


func (r *ResidenceRepo) GetResidence(req domain.GetResidenceRequest) (*domain.Residence, error) {
	residence := domain.Residence{}
	err := r.db.Table("residences").Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Take(&residence).Error // TODO
	if err != nil {
		fmt.Printf("get residence db error %v\n", err)
		return nil, err
	}
	return &residence, err
}

func (r *ResidenceRepo) GetResidenceList(req domain.GetResidenceListRequest) (*domain.MetaResidences, error) {
	var residences []domain.Residence
	whereQuery := "user_id_owner = ? AND deleted_at IS NULL AND "
	query := r.db.Debug().Table("residences")

	if len(req.LastId) > 0 && req.LastCreatedAt != nil {
		if req.Order == "desc" {
			query = query.
				Where(whereQuery+"(created_at < ? OR (id < ? AND created_at = ?))", req.UserIdOwner, req.LastCreatedAt, req.LastId, req.LastCreatedAt)
		} else {
			query = query.
				Where(whereQuery+"(created_at > ? OR (id > ? AND created_at = ?))", req.UserIdOwner, req.LastCreatedAt, req.LastId, req.LastCreatedAt)
		}
	} else {
		query = query.Where("user_id_owner = ? AND deleted_at IS NULL", req.UserIdOwner)
	}

	err := query.Order("created_at "+req.Order+", id "+req.Order).
		Limit(req.PerPage).
		Find(&residences).Error
	if err != nil {
		return nil, err
	}

	newLastId := ""
	var newLastCreatedAt *time.Time
	if len(residences) > 0 {
		newLastId = residences[len(residences)-1].Id
		newLastCreatedAt = residences[len(residences)-1].CreatedAt
	}

	return &domain.MetaResidences{
		Meta: domain.Meta{
			PerPage: req.PerPage,
			Order: req.Order,
			LastId: newLastId,
			LastCreatedAt: newLastCreatedAt,
		},
		Residences: residences,
	}, nil
}

func (r *ResidenceRepo) DeleteResidence(req domain.DeleteResidenceRequest) (error) {
	err := r.db.Table("residence").Where("id = ? and user_id_owner = ?", req.Id, req.UserIdOwner).Updates(map[string]interface{}{
		"deleted_at": time.Now(), // TODO
	}).Error
	if err != nil {
		fmt.Printf("delete residence db error %v\n", err)
		return err
	}
	return nil
}