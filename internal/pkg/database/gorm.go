package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
