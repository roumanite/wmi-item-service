package repository

import (
	"gorm.io/gorm"
	"wmi-item-service/internal/core/domain"
	"github.com/google/uuid"
	"errors"
	"fmt"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(req domain.CreateUserRequest) error {
	var users []domain.User
	err := r.db.Where("(email = ? OR username = ?) AND deleted_at IS NULL", req.Email, req.Username).Find(&users).Error
	if err != nil {
		return err
	}

	details := make(map[string]interface{})

	for _, user := range users {
		if user.Email == req.Email {
			details["email"] = map[string]interface{}{
				"type": "unique",
				"message": "Email is taken",
			}
		} else if user.Username == req.Username {
			details["username"] = map[string]interface{}{
				"type": "unique",
				"message": "Username is taken",
			}
		}
	}

	if len(details) > 0 {
		return domain.CustomError(domain.InvalidRequest, "", details)
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
		return domain.ErrUnknown
	}

	return nil
}

func ( r *UserRepo) GetUser(identifier string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Table("users").
		Where("email = ? OR username = ?", identifier, identifier).
		Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		fmt.Printf("get user db error %v\n", err)
		return nil, domain.ErrUnknown
	}
	return &user, nil
}