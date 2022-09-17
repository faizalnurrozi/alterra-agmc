package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTMiddleware struct {
}

func (jm JWTMiddleware) CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

/*func (jm JWTMiddleware) ExtractTokenUser(ctx echo.Context) bool {
	//a := jm.test("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjMzNDc1ODUsInVzZXJJZCI6NX0.I8Hgl8ZewOCCHODBce_foj2ywyCbtAprChLQHtGx7jM")
	validateToken := jm.validateToken()
	expiredTime := int64(validateToken.(jwt.MapClaims)["exp"].(float64))
	status := false
	if time.Now().Unix() < expiredTime {
		status = true
	}
	return status
}*/

func (jm JWTMiddleware) ValidateToken(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid 1")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid 2")
		}

		return jwt.SigningMethodHS256, nil
	})
	expiredTime := int64(token.Claims.(jwt.MapClaims)["exp"].(float64))
	status := false
	if time.Now().Unix() < expiredTime {
		status = true
	}

	return status
}
