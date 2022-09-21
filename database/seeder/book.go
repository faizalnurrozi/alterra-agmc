package seeder

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func bookSeeder(db *gorm.DB) {
	now := time.Now()
	var divisions = []model.Book{
		{
			Title:         "Title 1",
			Isbn:          "Isbn 1",
			Author:        "Author 1",
			Publisher:     "Publisher 1",
			DatePublished: now,
			StatusDisplay: false,
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Title:         "Title 2",
			Isbn:          "Isbn 2",
			Author:        "Author 2",
			Publisher:     "Publisher 2",
			DatePublished: now,
			StatusDisplay: false,
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Title:         "Title 3",
			Isbn:          "Isbn 3",
			Author:        "Author 3",
			Publisher:     "Publisher 3",
			DatePublished: now,
			StatusDisplay: false,
			Common: model.Common{
				ID:        3,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&divisions).Error; err != nil {
		log.Printf("cannot seed data books, with error %v\n", err)
	}
	log.Println("success seed data books")
}
