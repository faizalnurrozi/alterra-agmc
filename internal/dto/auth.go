package dto

import "github.com/golang-jwt/jwt/v4"

type (
	RegisterUserRequestBody struct {
		Name          string `json:"name" validate:"required"`
		Email         string `json:"email" validate:"required"`
		Password      string `json:"password" validate:"required"`
		Gender        string `json:"gender" validate:"required"`
		Nik           string `json:"nik" validate:"required"`
		BirthDate     string `json:"birth_date" validate:"required"`
		MarriedStatus bool   `json:"married_status" validate:"required"`
		YearOfJoin    int    `json:"year_of_join" validate:"required"`
	}

	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		UserID     uint   `json:"user_id"`
		Email      string `json:"email"`
		RoleID     uint   `json:"role_id"`
		DivisionID uint   `json:"division_id"`
		jwt.RegisteredClaims
	}
)
