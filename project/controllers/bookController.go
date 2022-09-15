package controllers

import (
	"day2-crud/config"
	"day2-crud/lib/database"
	"day2-crud/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type BookController struct {
	config.HTTPResponse
}

func NewBookController() *BookController {
	return &BookController{}
}

func (uc BookController) GetBooks(ctx echo.Context) error {
	var repositoryBook database.BookRepository
	books, err := repositoryBook.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, books)
}

func (uc BookController) GetBookByID(ctx echo.Context) error {
	var repositoryBook database.BookRepository
	book, err := repositoryBook.GetByID(2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, book)
}

func (uc BookController) Create(ctx echo.Context) error {
	var repositoryBook database.BookRepository
	book, err := repositoryBook.Create()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, book)
}

func (uc BookController) Update(ctx echo.Context) error {
	var repositoryBook database.BookRepository
	var data = models.Book{
		Title:         "Title updated",
		Isbn:          "ISBN updated",
		Author:        "Author updated",
		Publisher:     "Publisher updated",
		DatePublished: time.Now(),
		StatusDisplay: false,
	}
	book, err := repositoryBook.Update(data, 4)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, book)
}

func (uc BookController) Delete(ctx echo.Context) error {
	var repositoryBook database.BookRepository
	_, err := repositoryBook.Delete(4)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, nil)
}
