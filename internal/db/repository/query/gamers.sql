-- name: CreateGamer :one
INSERT OR REPLACE INTO gamers (
    challenge_id, user_id, start_date, end_date, target
) VALUES ( ?, ?, ?, ?, ?)
    RETURNING *;

-- name: DeleteGamer :exec
DELETE FROM gamers
WHERE user_id = ? and challenge_id  = ?;

-- name: ListGamers :many
SELECT
    challenge_id, user_id, name as challenge_name, user_name, start_date, end_date, target
FROM gamers AS l
JOIN register_users AS r ON l.user_id = r.id AND active = 1
JOIN challenges AS c ON l.challenge_id = c.id
WHERE challenge_id = ?;

-- name: GetLongestActivityPerDay :many
WITH max_distance_activities AS (
    SELECT
        a.user_id,
        strftime('%Y-%m-%d', datetime(a.start_date_local, 'unixepoch')) AS date,
        MAX(a.distance) AS max_distance
    FROM raw_activities AS a
             INNER JOIN gamers AS g ON a.user_id = g.user_id AND g.challenge_id = ?
    WHERE a.start_date_local BETWEEN g.start_date AND g.end_date
    GROUP BY a.user_id, date
)
SELECT a.*
FROM raw_activities AS a
 INNER JOIN max_distance_activities AS mda ON a.user_id = mda.user_id AND a.distance = mda.max_distance;

-- name: GetLongestActivityPerDay2 :many
WITH max_distance_activities AS (
SELECT
    a.user_id,
    strftime('%Y-%m-%d', datetime(a.start_date_local, 'unixepoch')) AS date,
    MAX(a.distance) AS max_distance
FROM raw_activities AS a
    INNER JOIN register_users AS g ON a.user_id = g.id
WHERE a.start_date_local >= sqlc.arg(begin) AND a.start_date_local <= sqlc.arg(end) AND a.sport_type='Run'
GROUP BY a.user_id, date
)
SELECT mda.date,
       a.id as activity_id,
       a.start_date_local,
       a.distance,
       a.average_speed,
       CAST((1000 / a.average_speed) / 60 AS REAL) AS pace,
       a.moving_time,
       a.name as activity_title,
       a.max_speed,
       g.first_name,
       g.last_name,
       g.user_name,
       g.profile,
       g.profile_medium
FROM raw_activities AS a
 INNER JOIN max_distance_activities AS mda ON a.user_id = mda.user_id AND a.distance = mda.max_distance
 INNER JOIN register_users AS g ON a.user_id = g.id
ORDER BY mda.date