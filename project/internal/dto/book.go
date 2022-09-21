package dto

import (
	"gorm.io/gorm"
	"time"
)

type (
	CreateBookRequestBody struct {
		Title         *string    `json:"title" validate:"required"`
		Isbn          *string    `json:"isbn" validate:"required"`
		Author        *string    `json:"author" validate:"required"`
		Publisher     *string    `json:"publisher" validate:"required"`
		DatePublished *time.Time `json:"date_published" validate:"required"`
		StatusDisplay *bool      `json:"status_display" validate:"required"`
	}
	UpdateBookRequestBody struct {
		ID            *uint      `param:"id" validate:"required"`
		Title         *string    `json:"title" validate:"required"`
		Isbn          *string    `json:"isbn" validate:"required"`
		Author        *string    `json:"author" validate:"required"`
		Publisher     *string    `json:"publisher" validate:"required"`
		DatePublished *time.Time `json:"date_published" validate:"required"`
		StatusDisplay *bool      `json:"status_display" validate:"required"`
	}
	BookResponse struct {
		ID            uint      `json:"id"`
		Title         string    `json:"title"`
		Isbn          string    `json:"isbn"`
		Author        string    `json:"author"`
		Publisher     string    `json:"publisher"`
		DatePublished time.Time `json:"date_published"`
		StatusDisplay bool      `json:"status_display"`
	}
	BookWithCUDResponse struct {
		BookResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
