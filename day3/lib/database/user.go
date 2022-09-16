package database

import (
	"mvcapp/config"
	"mvcapp/middlewares"
	"mvcapp/models"
)

func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func DetailUser(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.Where("id = ?", id).First(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func StorelUser(user models.User) error {
	return config.DB.Save(&user).Error
}

func UpdateUser(user models.User) error {
	return config.DB.Updates(&user).Error
}

func DeleteUser(id int) error {
	return config.DB.Delete(&models.User{}, id).Error
}
