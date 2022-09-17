package requests

import (
	"project/models"
	"time"
)

type UserRequest struct {
	Name          string               `json:"name" validate:"required"`
	Email         string               `json:"email" validate:"required,email"`
	Password      string               `json:"password"`
	Gender        models.GenderAllowed `json:"gender" sql:"type:gender" validate:"required"`
	Nik           string               `json:"nik" validate:"required"`
	BirthDate     time.Time            `json:"birth_date" validate:"required"`
	MarriedStatus bool                 `json:"married_status"`
	YearOfJoin    int                  `json:"year_of_join"`
}
