package domain

type Residence struct {
	Nickname string
	StreetAddress string
	City string
	State string
	Country string
	ZipCode string
	BuildingName string
}

type CreateResidenceRequest struct {
	UserID string
	Nickname string
	StreetAddress string
	City string
	State string
	Country string
	ZipCode string
	BuildingName string
}