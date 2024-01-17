package main

import (
	"fmt"
	"go-jwt-rbac/config"
	"go-jwt-rbac/controller"
	"go-jwt-rbac/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {
	conf := config.GetConfig()
	config.Connect(conf)

	route := gin.Default()

	routev1 := route.Group("/v1")
	routev1.POST("/register", controller.Register)
	routev1.POST("/login", controller.Login)
	routev1.GET("/api/category-products", controller.GetProductCategories)

	adminRoutev1 := routev1.Group("/admin")
	adminRoutev1.Use(middleware.JWTAdminAuth())
	adminRoutev1.POST("/category-products", controller.AddProductCategories)

	buyerRoutes := routev1.Group("/api")
	buyerRoutes.GET("/products", controller.GetAllProducts)
	buyerRoutes.GET("/products/{id}", controller.GetProductById)
	buyerRoutes.GET("/products/c/{id}", controller.GetProductByCategory)

	sellerRoutes := routev1.Group("/api")
	sellerRoutes.POST("/products", controller.AddProducts)

	route.Run(fmt.Sprintf(":%v", conf.ServerPort))
	fmt.Printf("Server starting at localhost:%v ...\n", conf.ServerPort)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading file env : %v", err)
	}

	StartServer()
}
