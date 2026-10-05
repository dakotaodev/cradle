package api

import (
	"context"
	"net/http"
	"time"

	"github.com/dakotaodev/cradle/internal/babies"
	"github.com/dakotaodev/cradle/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BabyRepository interface {
	CreateBaby(ctx context.Context, input db.CreateBabyParams) (babies.Baby, error)
	GetBaby(ctx context.Context, id string) (babies.Baby, error)
}

type CreateBabyRequest struct {
	FullName  string    `json:"fullName" binding:"required"`
	BirthDate time.Time `json:"birthDate"`
}

func CreateBabyHandler(repo BabyRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var json CreateBabyRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input := babies.CreateInput{
			FullName:  json.FullName,
			BirthDate: json.BirthDate,
		}

		baby, err := repo.CreateBaby(ctx, input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, baby)

	}
}

func GetBabyHandler(repo BabyRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		babyId := ctx.Param("babyId")

		_, err := uuid.Parse(babyId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		baby, err := repo.GetBaby(ctx, babyId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, baby)
	}
}
