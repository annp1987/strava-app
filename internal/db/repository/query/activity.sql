-- name: CreateActivity :one
INSERT OR IGNORE INTO raw_activities (
    id, user_id, create_at, start_date, distance, average_speed, moving_time, name, sport_type, max_speed, original_data
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    RETURNING id;