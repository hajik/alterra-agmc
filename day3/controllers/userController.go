package controllers

import (
	"errors"
	"mvcapp/lib/database"
	"mvcapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)

	users, err := database.LoginUser(&u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	res, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   res,
	})
}

func DetailUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, e := database.DetailUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   user,
	})
}

func StoreUserController(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if u == nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.New("Failed payload"))
	}

	err := database.StorelUser(*u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   u,
	})
}

func UpdateUserController(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if u == nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.New("Failed payload"))
	}

	e := database.UpdateUser(*u)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   u,
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
