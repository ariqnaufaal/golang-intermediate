package service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"GLIM_Hacktiv8/golang-intermediate/assignment-1/config"
	"GLIM_Hacktiv8/golang-intermediate/assignment-1/repository"
)

var DB *gorm.DB

var err error

func Connecting() {
	DB, err = gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}

	DB.AutoMigrate(&repository.Todo{})
}
