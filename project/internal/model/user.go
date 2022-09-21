package model

import (
	"time"
)

type GenderAllowed string

type User struct {
	Name          string    `json:"name" form:"name" validate:"required"`
	Email         string    `json:"email" form:"email" validate:"required,email"`
	Password      string    `json:"password" form:"password" validate:"required"`
	Gender        string    `json:"gender" form:"gender" sql:"type:gender" validate:"required"`
	Nik           string    `json:"nik" form:"nik" validate:"required"`
	BirthDate     time.Time `json:"birth_date" form:"birth_date" validate:"required"`
	MarriedStatus bool      `json:"married_status" form:"married_status"`
	YearOfJoin    int       `json:"year_of_join" form:"year_of_join"`
	Common
}
