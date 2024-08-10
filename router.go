package main

import (
	"net/http"
	"twigs-commerce-middleware/utils/middleware"

	"github.com/gin-gonic/gin"
)

func sample(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the twigs middleware!")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(request *gin.Context) {
		request.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.Use(middleware.CorsMiddleware())
	router.GET("/sample", sample)

	router.Use(middleware.AuthMiddelware())

	return router
}
