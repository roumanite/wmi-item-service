package service

import (
	"wmi-item-service/internal/core/port"
	"wmi-item-service/internal/core/domain"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetProfile(id string) (*domain.UserProfile, error) {
	profile, err := s.repo.GetProfile(id)
	return profile, err
}

func (s *UserService) UpdateProfile(req domain.UpdateProfileRequest) (*domain.UserProfile, error) {
	profile, err := s.repo.UpdateProfile(req)
	return profile, err
}