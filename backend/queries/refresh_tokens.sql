-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    id, user_id, token_hash, expires_at, created_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('user_id'),
    sqlc.arg('token_hash'),
    sqlc.arg('expires_at'),
    now()
) RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token_hash = sqlc.arg('token_hash')
  AND expires_at > now();

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token_hash = sqlc.arg('token_hash');

-- name: DeleteRefreshTokensByUser :exec
DELETE FROM refresh_tokens
WHERE user_id = sqlc.arg('user_id');
