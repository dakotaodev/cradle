-- name: CreateDiaperEvent :one
INSERT INTO diaper_events (
    baby_id, diaper_type, occurred_at, notes
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListDiaperEventsByBaby :many
SELECT * FROM diaper_events
WHERE baby_id = $1
ORDER BY occurred_at DESC, id DESC
LIMIT 20;