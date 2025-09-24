package database

import (
	"wishlistApi/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := "host=localhost user=user password=pass dbname=wishList port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Wish{})
	db.AutoMigrate(&model.WishList{})

	if !db.Migrator().HasConstraint(&model.WishList{}, "Wishes") {
		db.Migrator().CreateConstraint(&model.WishList{}, "Wishes")
	}

	if !db.Migrator().HasConstraint(&model.Wish{}, "fk_wish_list_wishes") {
		db.Migrator().CreateConstraint(&model.Wish{}, "fk_wish_list_wishes")
	}

	if !db.Migrator().HasConstraint(&model.WishList{}, "UserId") {
		db.Migrator().CreateConstraint(&model.WishList{}, "UserId")
	}

	if !db.Migrator().HasConstraint(&model.User{}, "fk_wish_list_user") {
		db.Migrator().CreateConstraint(&model.User{}, "fk_wish_list_user")
	}

}
