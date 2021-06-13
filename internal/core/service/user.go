package service

import (
	"wmi-item-service/internal/core/port"
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/core/service/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(req domain.CreateUserRequest) error {
	ep, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}
	err = s.repo.CreateUser(domain.InsertUserRequest{
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Username: req.Username,
		EncryptedPassword: ep,
	})
	return err
}
