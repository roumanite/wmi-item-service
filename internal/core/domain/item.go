package domain

type Item struct {
	Base
	Name string `json:"name"`
	UserIdOwner string `json:"userIdOwner"`
	CategoryId int `json:"categoryId"`
	DisplayPictureUrl string `json:"displayPictureUrl"` 
	Notes string `json:"notes"`
}

type CreateItemRequest struct {
	Name string
	UserIdOwner string
	CategoryId int
	DisplayPictureUrl string
	Notes string
}

type UpdateItemRequest struct {
	Id int
	Name string
	UserIdOwner string
	CategoryId int
	DisplayPictureUrl string
	Notes string
}

type GetItemRequest struct {
	Id int
	UserIdOwner string
}

type GetItemListRequest struct {
	UserIdOwner string
	PerPage int
	Order string
}

type DeleteItemRequest struct {
	Id int
	UserIdOwner string
}

type Meta struct {
	PerPage int `json:"perPage"`
	LastId string `json:"lastId"`
	Order string `json:"order"`
}

type MetaItems struct {
	Meta Meta `json:"meta"`
	Items []Item `json:"results"`
}
