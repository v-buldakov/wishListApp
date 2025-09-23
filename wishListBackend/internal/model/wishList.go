package model

type WishList struct {
	Id     int64
	Name   string `gorm:"not null;default:null"`
	Wishes []Wish `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId int64
}
