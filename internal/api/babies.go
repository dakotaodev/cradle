package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dakotaodev/cradle/internal/babies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BabyRepository interface {
	CreateBaby(ctx context.Context, input babies.CreateInput) (babies.Baby, error)
	GetBaby(ctx context.Context, id string) (babies.Baby, error)
}

type CreateBabyRequest struct {
	FullName  string `json:"fullName" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

type BabyResponse struct {
	ID string `json:"id" binding:"required"`
	FullName string `json:"fullName"`
	BirthDate string `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
}

func prepareResponse(baby babies.Baby) (BabyResponse) {
	return BabyResponse{
		ID: baby.ID,
		FullName: baby.FullName,
		BirthDate: baby.BirthDate.Format(time.DateOnly),
		CreatedAt: baby.CreatedAt,
	}
}

func CreateBabyHandler(repo BabyRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var json CreateBabyRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		birthDate, err := time.Parse(time.DateOnly, json.BirthDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := babies.CreateInput{
			FullName:  json.FullName,
			BirthDate: birthDate,
		}

		baby, err := repo.CreateBaby(ctx, input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, prepareResponse(baby))

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
			if errors.Is(err, babies.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, prepareResponse(baby))
	}
}
