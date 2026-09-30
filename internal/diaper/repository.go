package diaper

import (
	"context"

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

func (r *Repository) Create(ctx context.Context, input CreateInput) (Event, error) {

	params, err := toCreateParams(input)
	if err != nil {
		return Event{}, err
	}

	diaperEvent, err := r.db.CreateDiaperEvent(ctx, params)
	if err != nil {
		return Event{}, err
	}

	return convertDiaperEvent(diaperEvent)
}

func (r *Repository) ListRecent(ctx context.Context, babyId string) ([]Event, error) {

	parsedBabyId, err := uuid.Parse(babyId)
	if err != nil {
		return []Event{}, err
	}
	diaperEvents, err := r.db.ListDiaperEventsByBaby(ctx, pgtype.UUID{Bytes: parsedBabyId, Valid: true})

	if err != nil {
		return []Event{}, err
	}
	events := make([]Event, 0)
	for _, diaperEvent := range diaperEvents {
		event, err := convertDiaperEvent(diaperEvent)
		if err != nil {
			return []Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func convertDiaperEvent(diaperEvent db.DiaperEvent) (Event, error) {
	return Event{
		ID:         diaperEvent.ID.String(),
		BabyID:     diaperEvent.BabyID.String(),
		Type:       Type(diaperEvent.DiaperType),
		Notes:      diaperEvent.Notes,
		CreatedAt:  diaperEvent.CreatedAt.Time,
		OccurredAt: diaperEvent.OccurredAt.Time,
	}, nil
}
