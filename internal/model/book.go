package model

import "time"

type Book struct {
	Title         string    `json:"title" form:"title"`
	Isbn          string    `json:"isbn" form:"isbn"`
	Author        string    `json:"author" form:"author"`
	Publisher     string    `json:"publisher" form:"publisher"`
	DatePublished time.Time `json:"date_published" form:"date_published"`
	StatusDisplay bool      `json:"status_display" form:"status_display"`
	Common
}
