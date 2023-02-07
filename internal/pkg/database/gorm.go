package database

import (
	"fmt"
	"go_framework/internal/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDb(configs configs.Configs) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?&loc=Local&parseTime=True",
		configs.DB.Mysql.UserName,
		configs.DB.Mysql.Password,
		configs.DB.Mysql.Host,
		configs.DB.Mysql.Port,
		configs.DB.Mysql.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
