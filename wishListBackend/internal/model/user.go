package model

type User struct {
	Id        int64
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Token     string
	WishLists []WishList `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
