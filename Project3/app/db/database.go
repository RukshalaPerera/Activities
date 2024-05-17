package db

import (
	"Project3/app/Model"
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
	err = database.AutoMigrate(&model.User{}, &model.Book{}, &model.Reservation{})
	if err != nil {
		panic("failed to migrate database")
	}
	DB = database
}
