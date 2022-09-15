package database

import (
	"mvcapp/config"
	"mvcapp/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func DetailUser(id int) (interface{}, error) {
	var users models.User

	if e := config.DB.Where("id = ?", id).First(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func StorelUser(user models.User) error {
	return config.DB.Save(&user).Error
}

func UpdateUser(user models.User) error {
	return config.DB.Updates(&user).Error
}

func DeleteUser(id int) error {
	return config.DB.Delete("id", id).Error
}
