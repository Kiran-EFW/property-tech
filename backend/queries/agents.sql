-- name: CreateAgent :one
INSERT INTO agents (
    id, user_id, rera_number, pan, gst,
    tier, is_active, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('user_id'),
    sqlc.narg('rera_number'),
    sqlc.narg('pan'),
    sqlc.narg('gst'),
    COALESCE(sqlc.narg('tier')::varchar, 'bronze'),
    true,
    now(),
    now()
) RETURNING *;

-- name: GetAgentByID :one
SELECT * FROM agents
WHERE id = sqlc.arg('id');

-- name: GetAgentByUserID :one
SELECT * FROM agents
WHERE user_id = sqlc.arg('user_id');

-- name: ListAgents :many
SELECT * FROM agents
WHERE (sqlc.narg('tier')::varchar IS NULL OR tier = sqlc.narg('tier'))
  AND (sqlc.narg('is_active')::boolean IS NULL OR is_active = sqlc.narg('is_active'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateAgent :one
UPDATE agents SET
    rera_number = COALESCE(sqlc.narg('rera_number')::varchar, rera_number),
    pan         = COALESCE(sqlc.narg('pan')::varchar, pan),
    gst         = COALESCE(sqlc.narg('gst')::varchar, gst),
    tier        = COALESCE(sqlc.narg('tier')::varchar, tier),
    is_active   = COALESCE(sqlc.narg('is_active')::boolean, is_active),
    updated_at  = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpdateAgentTier :exec
UPDATE agents SET
    tier       = sqlc.arg('tier'),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: CountActiveAgents :one
SELECT count(*) FROM agents
WHERE is_active = true;
