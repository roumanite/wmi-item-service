package domain

type Residence struct {
	Base
	UserId string `json:"userId"`
	Nickname string `json:"nickname"`
	StreetAddress string `json:"streetAddress"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
	BuildingName string `json:"buildingName"`
}

type CreateResidenceRequest struct {
	UserId string
	Nickname string
	StreetAddress string
	City string
	State string
	Country string
	ZipCode string
	BuildingName string
}