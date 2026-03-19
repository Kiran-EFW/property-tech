-- name: CreateUser :one
INSERT INTO users (
    id, name, phone, email, password_hash, role, is_nri, avatar_url, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('name'),
    sqlc.arg('phone'),
    sqlc.narg('email'),
    sqlc.arg('password_hash'),
    sqlc.arg('role'),
    COALESCE(sqlc.narg('is_nri')::boolean, false),
    sqlc.narg('avatar_url'),
    now(),
    now()
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = sqlc.arg('id');

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = sqlc.arg('phone');

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = sqlc.arg('email');

-- name: UpdateUser :one
UPDATE users SET
    name         = COALESCE(sqlc.narg('name')::varchar, name),
    email        = COALESCE(sqlc.narg('email')::varchar, email),
    avatar_url   = COALESCE(sqlc.narg('avatar_url')::varchar, avatar_url),
    is_nri       = COALESCE(sqlc.narg('is_nri')::boolean, is_nri),
    updated_at   = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
WHERE (sqlc.narg('role')::varchar IS NULL OR role = sqlc.narg('role'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
