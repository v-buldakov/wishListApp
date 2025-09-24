package model

type WishList struct {
	Id     int32
	Name   string `gorm:"not null;"`
	Wishes []Wish `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId int64
}
