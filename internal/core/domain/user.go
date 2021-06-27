package domain

import "time"

type User struct {
	Id string
	Email string
	FirstName string
	LastName string
	Username string
	EncryptedPassword string
}

type UserProfile struct {
	Id string `json:"id"`
	Bio string `json:"bio"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Birthdate *time.Time `json:"birthdate"`
}

type SignUpRequest struct {
	Email string
	FirstName string
	LastName string
	Username string
	Password string
}

type SignInRequest struct {
	Identifier string
	Password string
}

type CreateUserRequest struct {
	Email string
	FirstName string
	LastName string
	Username string
	EncryptedPassword string
}

type UpdateProfileRequest struct {
	Id string
	Bio string
	FirstName string
	LastName string
	Birthdate *time.Time
}