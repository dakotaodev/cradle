package diaper

import (
	"github.com/dakotaodev/cradle/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db db.Querier
}

func NewRepository(database db.Querier) *Repository {
	return &Repository{db: database}
}

func toCreateParams(input CreateInput) (db.CreateDiaperEventParams, error) {
	// validate the input first
	if err := input.Validate(); err != nil {
		return db.CreateDiaperEventParams{}, err
	}

	babyId, err := uuid.Parse(input.BabyID)
	if err != nil {
		return db.CreateDiaperEventParams{}, err
	}

	return db.CreateDiaperEventParams{
		BabyID:     pgtype.UUID{Bytes: babyId, Valid: true},
		DiaperType: string(input.Type),
		OccurredAt: pgtype.Timestamptz{Time: input.OccurredAt, Valid: true},
		Notes:      input.Notes,
	}, nil
}
