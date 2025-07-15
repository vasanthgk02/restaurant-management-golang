package main

import (
	"net/http"
	"os"
	"restaurant_management/middleware"
	"restaurant_management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Server running successfully"})
	})

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)

	// Catch-all handler for undefined routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Route not found",
			"path":    c.Request.URL.Path,
			"message": "Check your endpoint",
		})
	})

	PORT := os.Getenv("PORT")
	router.Run(PORT)
}
