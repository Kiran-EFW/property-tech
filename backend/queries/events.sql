-- name: CreateEvent :one
INSERT INTO events (
    id, actor_id, actor_role, action,
    entity_type, entity_id, payload, created_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('actor_id'),
    sqlc.arg('actor_role'),
    sqlc.arg('action'),
    sqlc.arg('entity_type'),
    sqlc.arg('entity_id'),
    COALESCE(sqlc.narg('payload')::jsonb, '{}'::jsonb),
    now()
) RETURNING *;

-- name: ListEventsByEntity :many
SELECT * FROM events
WHERE entity_type = sqlc.arg('entity_type')
  AND entity_id = sqlc.arg('entity_id')
ORDER BY created_at DESC;

-- name: ListEventsByActor :many
SELECT * FROM events
WHERE actor_id = sqlc.arg('actor_id')
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListRecentEvents :many
SELECT * FROM events
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
