package seeder

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func userSeeder(db *gorm.DB) {
	now := time.Now()
	var divisions = []model.User{
		{
			Name:          "User 1",
			Email:         "user1@example.com",
			Password:      "$2a$14$yOtpt4Bq1JkPEkVNRBf.mOrGT.lL5oQ.Wl26K1e5I/AYcCHkEQZJS",
			Gender:        "MALE",
			Nik:           "1234",
			BirthDate:     now,
			MarriedStatus: false,
			YearOfJoin:    2012,
			Token:         "",
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name:          "User 2",
			Email:         "user2@example.com",
			Password:      "$2a$14$yOtpt4Bq1JkPEkVNRBf.mOrGT.lL5oQ.Wl26K1e5I/AYcCHkEQZJS",
			Gender:        "MALE",
			Nik:           "5678",
			BirthDate:     now,
			MarriedStatus: false,
			YearOfJoin:    2013,
			Token:         "",
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name:          "User 3",
			Email:         "user3@example.com",
			Password:      "$2a$14$yOtpt4Bq1JkPEkVNRBf.mOrGT.lL5oQ.Wl26K1e5I/AYcCHkEQZJS",
			Gender:        "MALE",
			Nik:           "4321",
			BirthDate:     now,
			MarriedStatus: false,
			YearOfJoin:    2014,
			Token:         "",
			Common: model.Common{
				ID:        3,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&divisions).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}
