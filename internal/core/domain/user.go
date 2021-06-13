package domain

type User struct {
	Id string
	Email string
	FirstName string
	LastName string
	Username string
	EncryptedPassword string
}

type CreateUserRequest struct {
	Email string
	FirstName string
	LastName string
	Username string
	Password string
}

type InsertUserRequest struct {
	Email string
	FirstName string
	LastName string
	Username string
	EncryptedPassword string
}