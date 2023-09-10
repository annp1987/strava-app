-- name: CreateGamer :one
INSERT OR REPLACE INTO gamers (
    user_id, start_date, end_date, target
) VALUES ( ?, ?, ?, ?)
    RETURNING user_id;

-- name: DeleteGamer :exec
DELETE FROM gamers
WHERE user_id = ?;

-- name: ListGamers :many
SELECT
    user_id, start_date, end_date, target
FROM gamers;

-- name: ListLongestRunPerDay :many
SELECT
    l.user_id, today, activity_id, l.start_date, distance, average_speed, moving_time, name, sport_type, max_speed
FROM longest_run_per_day as l
INNER JOIN gamers AS r on l.user_id = r.user_id;

-- name: UpdateLongestRunPerDay :one
INSERT OR REPLACE into longest_run_per_day (
    user_id, today, activity_id, start_date, distance, average_speed, moving_time, name, sport_type, max_speed)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING user_id, today;

-- name: GetCurrentLongestRunPerDay :one
SELECT * FROM longest_run_per_day
WHERE user_id = ? AND today = ?