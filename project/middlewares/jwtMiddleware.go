package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

type JWTMiddleware struct {
}

func (jm JWTMiddleware) CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func (jm JWTMiddleware) ExtractTokenUser(ctx echo.Context) int {
	user := ctx.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}
