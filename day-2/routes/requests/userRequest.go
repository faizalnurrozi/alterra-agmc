package requests

import (
	"day2-crud/models"
	"time"
)

type UserRequest struct {
	Name          string               `json:"name"`
	Email         string               `json:"email"`
	Gender        models.GenderAllowed `json:"gender" sql:"type:gender"`
	Nik           string               `json:"nik"`
	BirthDate     time.Time            `json:"birth_date"`
	MarriedStatus bool                 `json:"married_status"`
	YearOfJoin    int                  `json:"year_of_join"`
}
