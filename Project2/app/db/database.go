package db

import (
	"Project2/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := "root:1234@tcp(localhost:3306)/fiberdb?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = database.AutoMigrate(&model.User{}, &model.Role{})
	if err != nil {
		panic("failed to migrate database")
	}
	DB = database
}
