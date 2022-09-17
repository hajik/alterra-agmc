package controllers

import (
	"errors"
	"mvcapp/lib/database"
	"mvcapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBookController(c echo.Context) error {
	books, err := database.GetBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   books,
	})
}

func DetailBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := database.DetailBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   book,
	})
}

func StoreBookController(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if book == nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.New("Failed payload"))
	}

	if err := database.StoreBook(*book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   book,
	})
}

func UpdateBookController(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if book == nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.New("Failed payload"))
	}

	if err := database.UpdateBook(*book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   book,
	})
}

func DeleteBookController(c echo.Context) error {
	if err := database.DeleteBook(1); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Book has been successfully deleted",
	})
}
