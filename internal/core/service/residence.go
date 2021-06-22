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

func (s *ResidenceService) CreateResidence(req domain.CreateResidenceRequest) (*domain.Residence, error) {
	residence, err := s.repo.CreateResidence(req)
	if err != nil {
		return nil, err
	}
	return residence, nil
}

// check *****
func (s *ResidenceService) UpdateResidence(req domain.UpdateResidenceRequest) (*domain.Residence, error) {
	residence, err := s.repo.UpdateResidence(req)
	if err != nil {
		return nil, err
	}
	return residence, nil
}

func (s *ResidenceService) GetResidence(req domain.GetResidenceRequest) (*domain.Residence, error) {
	residence, err := s.repo.GetResidence(req)
	if err != nil {
		return nil, err
	}
	return residence, nil
}

func (s *ResidenceService) GetResidenceList(req domain.GetResidenceListRequest) (*domain.MetaResidences, error) {
	residences, err := s.repo.GetResidenceList(req)
	if err != nil {
		return nil, err
	}
	return residences, nil
}

func (s *ResidenceService) DeleteResidence(req domain.DeleteResidenceRequest) (error) {
	err := s.repo.DeleteResidence(req)
	if err != nil {
		return err
	}
	return nil
}