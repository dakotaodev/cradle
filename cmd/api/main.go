package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func main() {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", health)

	router.Run(":8080")

}
