package database

import (
	"gorm.io/gorm"
	"project/config"
	"project/models"
	"time"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{DB: DB}
}

func (s BookRepository) GetAll() (interface{}, error) {
	var books []models.Book

	if err := s.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (s BookRepository) GetByID(id int) (interface{}, error) {
	var book models.Book

	if err := s.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (s BookRepository) Create() (interface{}, error) {
	book := models.Book{
		Title:         "Title create new",
		Isbn:          "ISBN create new",
		Author:        "Author create new",
		Publisher:     "Publisher create new",
		DatePublished: time.Now(),
		StatusDisplay: true,
	}

	if err := s.DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (s BookRepository) Update(data models.Book, id int) (interface{}, error) {
	var book models.Book

	if err := s.DB.Find(&book, id).Error; err != nil {
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

func (s BookRepository) Delete(id int) (interface{}, error) {
	var book models.Book

	if err := s.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}
