package controllers

import (
	"mvcapp/lib/database"
	"mvcapp/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	_, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	res := models.User{
		Name:     "Jojo si Jojo",
		Email:    "jojo@gmail.com",
		Password: "jojo123",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{

		"status": "success",
		"code":   http.StatusOK,
		"data": []models.User{
			res,
		},
	})
}

func DetailUserController(c echo.Context) error {
	_, e := database.DetailUser(1)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data": models.User{
			Name:     "Jojo si Jojo",
			Email:    "jojo@gmail.com",
			Password: "jojo123",
		},
	})
}

func StoreUserController(c echo.Context) error {
	e := database.StorelUser(models.User{})
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data": models.User{
			Name:     "Jojo si Jojo",
			Email:    "jojo@gmail.com",
			Password: "jojo123",
		},
	})
}

func UpdateUserController(c echo.Context) error {
	e := database.UpdateUser(models.User{})
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data": models.User{
			Name:     "Jojo si Jojo",
			Email:    "jojo@gmail.com",
			Password: "jojo123",
		},
	})
}

func DeleteUserController(c echo.Context) error {
	e := database.DeleteUser(1)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "User has been successfully deleted",
	})
}
