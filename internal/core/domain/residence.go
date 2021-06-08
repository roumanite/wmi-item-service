package domain

type Residence struct {
	Base
	UserIdOwner string `json:"UserIdOwner"`
	Nickname string `json:"nickname"`
	StreetAddress string `json:"streetAddress"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
	BuildingName string `json:"buildingName"`
}

type CreateResidenceRequest struct {
	UserIdOwner string
	Nickname string
	StreetAddress string
	City string
	State string
	Country string
	ZipCode string
	BuildingName string
}

// check *****
type UpdateResidenceRequest struct {
	Id int
	UserIdOwner string
	Nickname string
	StreetAddress string
	City string
	State string
	Country string
	ZipCode string
	BuildingName string
}

type GetResidenceRequest struct {
	Id int
	UserIdOwner string
}

type DeleteResidenceRequest struct {
	Id int
	UserIdOwner string
}
