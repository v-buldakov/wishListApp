package model

import "time"

type Wish struct {
	Id         int64
	Name       string `gorm:"not null;"`
	Url        string
	PicUrl     string
	WishListId int `gorm:"not null"`
	ComletedAt time.Time
}
