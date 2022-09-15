package config

import (
	"fmt"
	"mvcapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"db_username": "root",
		"db_password": "MyPassword@123",
		"db_port":     "3306",
		"db_host":     "localhost",
		"db_name":     "mb_altera",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config["db_username"], config["db_password"], config["db_host"], config["db_port"], config["db_name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
