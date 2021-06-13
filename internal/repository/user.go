package repository

import (
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
	"errors"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(req domain.InsertUserRequest) error {
	var users []domain.User
	err := r.db.Where("(email = ? OR username = ?) AND deleted_at IS NULL", req.Email, req.Username).Find(&users).Error
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Email == req.Email {
			return errors.New("duplicate email")
		} else if user.Username == req.Username {
			return errors.New("duplicate username")
		}
	}

	err = r.db.Create(&domain.User{
		Id: uuid.New().String(),
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Username: req.Username,
		EncryptedPassword: req.EncryptedPassword,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
