package babies

import (
	"context"

	"github.com/dakotaodev/cradle/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BabyQuerier interface {
	CreateBaby(ctx context.Context, arg db.CreateBabyParams) (db.Baby, error)
	GetBaby(ctx context.Context, pgtype.UUID string) (db.Baby, error)
}

type Repository struct {
	db BabyQuerier
}

func NewRepository(database BabyQuerier) *Repository {
	return &Repository{db: database}
}

func (r *Repository) CreateBaby(ctx context.Context, arg db.CreateBabyParams) (db.Baby, error) {
	params, err := toCreateParams(input)
	if err != nil {
		return Baby{}, nil
	}

	baby, err := r.db.CreateBaby(ctx, params)
	if err != nil {
		return Baby{}, nil
	}

	return baby, nil
}

func toCreateParams(input CreateInput) (db.CreateBabyParams, error) {
	if err := input.Validate(); err != nil {
		return db.CreateBabyParams{}, nil
	}

	return db.CreateBabyParams{
		FullName:  input.FullName,
		BirthDate: pgtype.Date{Time: input.BirthDate, Valid: true},
	}, nil
}

func (r *Repository) GetBaby(ctx context.Context, id pgtype.UUID) (db.Baby, error) {
	parseId, err := uuid.Parse(babyId)
	if err != nil {
		return db.Baby{}, err
	}

	baby, err := r.db.GetBaby(ctx, pgtype.UUID{Bytes: parseId, Valid: true})

	return baby, nil
}
