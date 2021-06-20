package service

import (
	"errors"
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

func (s *AuthService) SignIn(req domain.SignInRequest) (*domain.User, error) {
	user, err := s.repo.GetUser(req.Identifier)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidLoginDetails
		}
		return nil, err
	}

	if !util.CheckPasswordHash(req.Password, user.EncryptedPassword) {
		return nil, domain.ErrInvalidLoginDetails
	}

	return user, nil
}
