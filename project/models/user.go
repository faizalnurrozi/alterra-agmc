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
	Name          string        `json:"name" form:"name" validate:"required"`
	Email         string        `json:"email" form:"email" validate:"required,email"`
	Password      string        `json:"password" form:"password" validate:"required"`
	Gender        GenderAllowed `json:"gender" form:"gender" sql:"type:gender" validate:"required"`
	Nik           string        `json:"nik" form:"nik" validate:"required"`
	BirthDate     time.Time     `json:"birth_date" form:"birth_date" validate:"required"`
	MarriedStatus bool          `json:"married_status" form:"married_status"`
	YearOfJoin    int           `json:"year_of_join" form:"year_of_join"`
	Token         string        `json:"token" form:"token"`
}
