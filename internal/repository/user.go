package repository

import "wmi-item-service/pkg/postgresql"

type UserRepo struct {
	db *postgresql.Connection
}

func NewUserRepo(db *postgresql.Connection) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser() error {
	// TODO
	return nil
}