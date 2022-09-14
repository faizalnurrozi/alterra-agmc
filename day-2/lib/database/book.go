package database

import (
	"day2-crud/config"
	"day2-crud/models"
	"time"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(id int) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func StoreBook() (interface{}, error) {
	book := models.Book{
		Title:         "Title create new",
		Isbn:          "ISBN create new",
		Author:        "Author create new",
		Publisher:     "Publisher create new",
		DatePublished: time.Now(),
		StatusDisplay: true,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(data models.Book, id int) (interface{}, error) {
	var book models.Book

	if err := config.DB.Find(&book, id).Error; err != nil {
		return nil, err
	}

	book.Title = data.Title
	book.Isbn = data.Isbn
	book.Author = data.Author
	book.Publisher = data.Publisher
	book.DatePublished = data.DatePublished
	book.StatusDisplay = data.StatusDisplay
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}
