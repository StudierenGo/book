package main

import "github.com/gin-gonic/gin"

func main() {
	// ОБъект типа *gin.Engine
	router := gin.Default()

	router.POST("/employee")
	router.GET("/employee/:id")
	router.PUT("/employee/:id")
	router.DELETE("employee/:id")

	router.Run(":8080")
}
