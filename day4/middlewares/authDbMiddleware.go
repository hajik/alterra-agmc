package middlewares

import (
	"mvcapp/config"
	"mvcapp/models"

	"github.com/labstack/echo"
)

func BasiAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email=? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
