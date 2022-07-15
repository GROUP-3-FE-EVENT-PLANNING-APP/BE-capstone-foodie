package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SecretJWT")),
	})
}

func CreateToken(userId int, avatarUrl, role, handphone, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["avatarUrl"] = avatarUrl
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 504).Unix()
	claims["handphone"] = handphone
	claims["email"] = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SecretJWT")))
}

func ExtractToken(e echo.Context) (data map[string]interface{}, err error) {
	user := e.Get("user").(*jwt.Token)
	var dataToken = map[string]interface{}{}
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		dataToken["userId"] = claims["userId"].(float64)
		dataToken["avatarUrl"] = claims["avatarUrl"].(string)
		dataToken["role"] = claims["role"].(string)
		dataToken["handphone"] = claims["handphone"].(string)
		dataToken["email"] = claims["email"].(string)
		return dataToken, nil
	}
	return nil, fmt.Errorf("token invalid")
}
