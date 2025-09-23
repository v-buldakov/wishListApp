package model

type Wish struct {
	Id         int64
	Name       string `gorm:"not null;"`
	Url        string
	PicUrl     string
	WishListId int  `gorm:"not null"`
	Comleted   bool `gorm:"not null"`
}
