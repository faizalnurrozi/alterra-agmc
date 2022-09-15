package models

import (
	"gorm.io/gorm"
	"time"
)

type GenderAllowed string

const (
	MALE   GenderAllowed = "MALE"
	FEMALE GenderAllowed = "FEMALE"
)

type User struct {
	gorm.Model
	Name          string        `json:"name" form:"name"`
	Email         string        `json:"email" form:"email"`
	Password      string        `json:"password" form:"password"`
	Gender        GenderAllowed `json:"gender" form:"gender" sql:"type:gender"`
	Nik           string        `json:"nik" form:"nik"`
	BirthDate     time.Time     `json:"birth_date" form:"birth_date"`
	MarriedStatus bool          `json:"married_status" form:"married_status"`
	YearOfJoin    int           `json:"year_of_join" form:"year_of_join"`
	Token         string        `json:"token" form:"token"`
}
