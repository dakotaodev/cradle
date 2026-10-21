package babies

import (
	"context"
	"errors"

	"github.com/dakotaodev/cradle/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type BabyQuerier interface {
	CreateBaby(ctx context.Context, arg db.CreateBabyParams) (db.Baby, error)
	GetBaby(ctx context.Context, id pgtype.UUID) (db.Baby, error)
}

type Repository struct {
	db BabyQuerier
}

func NewRepository(database BabyQuerier) *Repository {
	return &Repository{db: database}
}

func (r *Repository) CreateBaby(ctx context.Context, input CreateInput) (Baby, error) {
	params, err := toCreateParams(input)
	if err != nil {
		return Baby{}, err
	}

	baby, err := r.db.CreateBaby(ctx, params)
	if err != nil {
		return Baby{}, err
	}

	return Baby{
		ID:        baby.ID.String(),
		FullName:  baby.FullName,
		BirthDate: baby.BirthDate.Time,
		CreatedAt: baby.CreatedAt.Time,
	}, nil
}

func toCreateParams(input CreateInput) (db.CreateBabyParams, error) {
	if err := input.Validate(); err != nil {
		return db.CreateBabyParams{}, err
	}

	return db.CreateBabyParams{
		FullName:  input.FullName,
		BirthDate: pgtype.Date{Time: input.BirthDate, Valid: true},
	}, nil
}

func (r *Repository) GetBaby(ctx context.Context, babyId string) (Baby, error) {
	parseId, err := uuid.Parse(babyId)
	if err != nil {
		return Baby{}, err
	}

	baby, err := r.db.GetBaby(ctx, pgtype.UUID{Bytes: parseId, Valid: true})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Baby{}, ErrNotFound
		}
		return Baby{}, err
	}
	return Baby{
		ID:        baby.ID.String(),
		FullName:  baby.FullName,
		BirthDate: baby.BirthDate.Time,
		CreatedAt: baby.CreatedAt.Time,
	}, nil
}
