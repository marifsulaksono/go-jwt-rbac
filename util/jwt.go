package util

import (
	"errors"
	"fmt"
	"go-jwt-rbac/model"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

	return token.SignedString([]byte(os.Getenv(secretKey)))
}

func GetTokenFromHeader(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return nil, errors.New("authentication required")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method invalid")
		}

		return []byte(os.Getenv(secretKey)), nil
	})

	return token, nil
}

func ValidateAdminToken(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if ok && token.Valid && role == "admin" {
		return nil
	}

	return errors.New("the token provided is not admin")
}

func ValidateSellerToken(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	fmt.Println(claims)
	if ok && token.Valid && role == "seller" {
		return nil
	}

	return errors.New("the token provided is not seller")
}

func ValidateBuyerToken(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if ok && token.Valid && role == "buyer" {
		return nil
	}

	return errors.New("the token provided is not buyer")
}
