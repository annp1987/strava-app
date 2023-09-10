-- name: GetToken :one
SELECT access_token, refresh_token, expired_at FROM register_users
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT OR REPLACE INTO register_users (
    id, user_name, profile_medium, profile, access_token, refresh_token, expired_at
) VALUES ( ?, ?, ?, ?, ?, ?, ?)
    RETURNING id;

-- name: UpdateUser :one
UPDATE register_users
set active = ?
WHERE id = ?
    RETURNING id;

-- name: UpdateToken :one
UPDATE register_users
set access_token = ?, refresh_token = ?, expired_at = ?
WHERE id = ?
    RETURNING id;