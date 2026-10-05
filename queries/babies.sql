-- name: CreateBaby :one
INSERT INTO babies(
    full_name,
    birth_date
)
VALUES(
    $1,
    $2
)
RETURNING *; 

-- name: GetBaby :one
SELECT * FROM babies
WHERE ID = $1
LIMIT 1;