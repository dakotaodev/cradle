package api

import (
	"github.com/dakotaodev/cradle/internal/diaper"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DiaperRequest struct {
	DiaperType string    `form:"diaperType" json:"diaperType" binding:"required"`
	Notes      string    `form:"notes" json:"notes"`
	OccurredAt time.Time `form:"occurredAt" json:"occurredAt"`
}

func DiaperHandler(c *gin.Context) {

	var json DiaperRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	babyId := c.Param("babyId")

	input := diaper.CreateInput{
		BabyID:     babyId,
		Type:       diaper.Type(json.DiaperType),
		Notes:      json.Notes,
		OccurredAt: json.OccurredAt,
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "ok"})
}
