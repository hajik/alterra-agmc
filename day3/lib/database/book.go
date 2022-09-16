package database

import (
	"mvcapp/config"
	"mvcapp/models"
)

func GetBook() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func DetailBook(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.Where("id = ?", id).First(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func StoreBook(book models.Book) error {
	return config.DB.Save(&book).Error
}

func UpdateBook(book models.Book) error {
	return config.DB.Updates(&book).Error
}

func DeleteBook(id int) error {
	return config.DB.Delete(&models.Book{}, id).Error
}
