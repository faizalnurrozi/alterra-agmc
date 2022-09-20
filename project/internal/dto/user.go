package dto

import (
	"gorm.io/gorm"
	"time"
)

type (
	UpdateUserRequestBody struct {
		ID            *uint   `param:"id" validate:"required"`
		Name          *string `json:"name" validate:"required"`
		Email         *string `json:"email" validate:"required"`
		Password      *string `json:"password" validate:"required"`
		Gender        *string `json:"gender" validate:"required"`
		Nik           *string `json:"nik" validate:"required"`
		BirthDate     *string `json:"birth_date" validate:"required"`
		MarriedStatus *bool   `json:"married_status" validate:"required"`
		YearOfJoin    *int    `json:"year_of_join" validate:"required"`
	}
	UserResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	UserWithJWTResponse struct {
		UserResponse
		JWT string `json:"jwt"`
	}
	UserWithCUDResponse struct {
		UserResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
