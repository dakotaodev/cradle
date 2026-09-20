-- name: CreateDiaperEvent :one
INSERT INTO diaper_events (
    baby_id, diaper_type, occurred_at, created_at, notes
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;