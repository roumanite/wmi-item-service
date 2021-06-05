package domain

type Item struct {
	Base
	Id int `json:"id"`
    Name string `json:"name"`
    UserIdOwner string `json:"userIdOwner"`
    CategoryId int `json:"categoryId"`
    DisplayPictureUrl string `json:"displayPictureUrl"` 
    Notes string `json:"notes"`
}

type CreateItemRequest struct {
	Base
	Id int
    Name string
    UserIdOwner string
    CategoryId int
    DisplayPictureUrl string
    Notes string
}