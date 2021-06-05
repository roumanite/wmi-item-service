package service

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/core/port"
)

type ResidenceService struct {
	repo port.ResidenceRepository
}

func NewResidenceService(repo port.ResidenceRepository) *ResidenceService {
	return &ResidenceService{repo}
}

func (s *ResidenceService) CreateResidence(req domain.CreateResidenceRequest) error {
	err := s.repo.CreateResidence(req)
	if err != nil {
		return err
	}
	return nil
}