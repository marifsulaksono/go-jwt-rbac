package util

import (
	"fmt"
	"go-jwt-rbac/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "JWT_SECRET_KEY"

func GenerateJWT(user model.UserResponse) (string, error) {
	jwtExpiredTime := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"RegisteredClaims": jwt.RegisteredClaims{
			Issuer:    "go-jwt-rbac",
			ExpiresAt: jwt.NewNumericDate(jwtExpiredTime),
		},
	})

	fmt.Println(token)

	return token.SignedString([]byte(os.Getenv(secretKey)))
}
