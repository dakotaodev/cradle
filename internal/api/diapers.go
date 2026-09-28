package api

import (
	"context"
	"net/http"
	"time"

	"github.com/dakotaodev/cradle/internal/diaper"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Create(ctx context.Context, input diaper.CreateInput) (diaper.Event, error)
}

type DiaperRequest struct {
	DiaperType string    `form:"diaperType" json:"diaperType" binding:"required"`
	Notes      string    `form:"notes" json:"notes"`
	OccurredAt time.Time `form:"occurredAt" json:"occurredAt"`
}

func DiaperHandler(repo Repository) gin.HandlerFunc {

	return func(c *gin.Context) {

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

		event, err := repo.Create(c.Request.Context(), input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, event)
	}

}
