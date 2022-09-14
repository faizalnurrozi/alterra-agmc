package controllers

import (
	"day2-crud/lib/database"
	"day2-crud/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetBooksController(ctx echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIDController(ctx echo.Context) error {
	book, err := database.GetBookByID(2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func StoreBook(ctx echo.Context) error {
	book, err := database.StoreBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func UpdateBook(ctx echo.Context) error {
	var data = models.Book{
		Title:         "Title updated",
		Isbn:          "ISBN updated",
		Author:        "Author updated",
		Publisher:     "Publisher updated",
		DatePublished: time.Now(),
		StatusDisplay: false,
	}
	book, err := database.UpdateBook(data, 4)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func DeleteBook(ctx echo.Context) error {
	_, err := database.DeleteBook(4)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   nil,
	})
}
