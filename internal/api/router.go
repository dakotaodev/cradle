package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(repo Repository, babyRepo BabyRepository) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", HealthHandler)
	router.POST("/babies/:babyId/diapers", DiaperHandler(repo))
	router.GET("/babies/:babyId/diapers", ListRecentHandler(repo))
	router.GET("/babies/:babyId", GetBabyHandler(babyRepo))
	router.POST("/babies/:babyId", CreateBabyHandler((babyRepo)))
	return router
}
