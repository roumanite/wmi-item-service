package domain

import "time"

type Base struct {
	Id        string `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

