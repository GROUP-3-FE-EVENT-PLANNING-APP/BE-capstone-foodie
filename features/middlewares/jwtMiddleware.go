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

func ExtractToken(e echo.Context) (int, string, string, string, string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		avatarUrl := claims["avatarUrl"].(string)
		role := claims["role"].(string)
		handphone := claims["handphone"].(string)
		email := claims["email"].(string)
		return int(userId), avatarUrl, role, handphone, email, nil
	}
	return 0, "Avatar link not found", "you dont have any role", "you dont have any number stored", "no email registered", fmt.Errorf("token invalid")
}
