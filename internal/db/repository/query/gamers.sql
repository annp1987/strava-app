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
    challenge_id, user_id, user_name, start_date, end_date, target
FROM gamers AS l
JOIN register_users AS r ON l.user_id = r.id AND active = 1
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
 INNER JOIN max_distance_activities AS mda ON a.user_id = mda.user_id AND a.distance = mda.max_distance