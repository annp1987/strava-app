-- name: CreateActivity :exec
INSERT OR IGNORE INTO raw_activities (
    id, user_id, create_at, start_date, start_date_local, distance, average_speed, moving_time, name, sport_type, max_speed, original_data
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetActivity :many
SELECT
    id,
    user_id,
    create_at,
    start_date,
    start_date_local,
    distance,
    average_speed,
    moving_time,
    name,
    sport_type,
    max_speed
FROM raw_activities
WHERE user_id=? AND sport_type=?;