package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/config"
	"project/lib/database"
	"project/models"
	"time"
)

type BookController struct {
	config.Contract
}

func NewBookController(contract config.Contract) *BookController {
	return &BookController{Contract: contract}
}

func (uc BookController) GetBooks(ctx echo.Context) error {
	if ctx.Request().Method == http.MethodGet {
		repositoryBook := database.NewBookRepository(uc.DB)
		books, err := repositoryBook.GetAll()
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}
		return uc.ResponseOk(ctx, http.StatusOK, books)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc BookController) GetBookByID(ctx echo.Context) error {
	if ctx.Request().Method == http.MethodGet {
		repositoryBook := database.NewBookRepository(uc.DB)
		book, err := repositoryBook.GetByID(2)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, book)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc BookController) Create(ctx echo.Context) error {
	if ctx.Request().Method == http.MethodPost && ctx.Request().Header.Get("Authorization") != "" {
		repositoryBook := database.NewBookRepository(uc.DB)
		book, err := repositoryBook.Create()
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusCreated, book)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc BookController) Update(ctx echo.Context) error {
	if ctx.Request().Method == http.MethodPut && ctx.Request().Header.Get("Authorization") != "" {
		repositoryBook := database.NewBookRepository(uc.DB)
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
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, book)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc BookController) Delete(ctx echo.Context) error {
	if ctx.Request().Method == http.MethodDelete && ctx.Request().Header.Get("Authorization") != "" {
		repositoryBook := database.NewBookRepository(uc.DB)
		_, err := repositoryBook.Delete(4)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, nil)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}
