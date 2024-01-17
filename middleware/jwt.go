package middleware

import (
	"fmt"
	"go-jwt-rbac/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.GetTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		err = util.ValidateAdminToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTSellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.GetTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		err = util.ValidateSellerToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		fmt.Println(token.Claims.(jwt.MapClaims)["id"])
		c.Set("user_id", token.Claims.(jwt.MapClaims)["id"])
		c.Next()
	}
}

func JWTBuyerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.GetTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		err = util.ValidateBuyerToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		fmt.Println(token.Claims.(jwt.MapClaims)["id"])
		c.Set("user_id", token.Claims.(jwt.MapClaims)["id"])
		c.Next()
	}
}
