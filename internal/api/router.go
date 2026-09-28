package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(repo Repository) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", HealthHandler)
	router.POST("/babies/:babyId/diapers", DiaperHandler(repo))

	return router
}
