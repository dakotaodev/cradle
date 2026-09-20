-- name: CreateDiaperEvent :one
INSERT INTO diaper_events (
    baby_id, diaper_type, occurred_at, notes
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;