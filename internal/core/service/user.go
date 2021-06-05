package service

import (
	"wmi-item-service/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser() error {
	// TODO
	return nil
}