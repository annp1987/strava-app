-- name: CreateChallenge :one
INSERT INTO challenges (
    name, rules
) VALUES ( ?, ?)
RETURNING *;

-- name: UpdateChallenge :one
UPDATE challenges
 SET name = ?, rules = ?
WHERE id = ?
RETURNING *;

-- name: GetChallenge :one
SELECT
    *
FROM challenges
WHERE id = ?;

-- name: ListChallenge :many
SELECT * FROM challenges;