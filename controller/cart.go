package controller

import (
	"go-jwt-rbac/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCartByUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetString("user_id"))
	carts, err := model.GetProductById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": carts})
}

func AddCart(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetString("user_id"))
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart.UserId = userId
	err := cart.CreateCart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success create cart"})
}
