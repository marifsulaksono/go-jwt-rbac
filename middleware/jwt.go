package middleware

import (
	"go-jwt-rbac/util"
	"net/http"

	"github.com/gin-gonic/gin"
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
