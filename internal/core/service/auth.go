package service

import (
	"wmi-item-service/internal/core/port"
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/core/service/util"
)

type AuthService struct {
	repo port.UserRepository
}

func NewAuthService(repo port.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) SignUp(req domain.SignUpRequest) error {
	ep, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}
	err = s.repo.CreateUser(domain.CreateUserRequest{
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Username: req.Username,
		EncryptedPassword: ep,
	})
	return err
}

func (s *AuthService) SignIn(req domain.SignInRequest) error {
	return nil
}
