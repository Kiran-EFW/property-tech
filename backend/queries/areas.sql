-- name: CreateArea :one
INSERT INTO areas (
    id, name, slug, city, state, description,
    location, price_trend, infrastructure,
    created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('name'),
    sqlc.arg('slug'),
    sqlc.arg('city'),
    sqlc.arg('state'),
    sqlc.narg('description'),
    ST_MakePoint(sqlc.arg('lng')::float8, sqlc.arg('lat')::float8)::geography,
    COALESCE(sqlc.narg('price_trend')::jsonb, '[]'::jsonb),
    COALESCE(sqlc.narg('infrastructure')::jsonb, '{}'::jsonb),
    now(),
    now()
) RETURNING *;

-- name: GetAreaByID :one
SELECT * FROM areas
WHERE id = sqlc.arg('id');

-- name: GetAreaBySlug :one
SELECT * FROM areas
WHERE slug = sqlc.arg('slug');

-- name: ListAreas :many
SELECT * FROM areas
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateArea :one
UPDATE areas SET
    name           = COALESCE(sqlc.narg('name')::varchar, name),
    slug           = COALESCE(sqlc.narg('slug')::varchar, slug),
    city           = COALESCE(sqlc.narg('city')::varchar, city),
    state          = COALESCE(sqlc.narg('state')::varchar, state),
    description    = COALESCE(sqlc.narg('description')::text, description),
    price_trend    = COALESCE(sqlc.narg('price_trend')::jsonb, price_trend),
    infrastructure = COALESCE(sqlc.narg('infrastructure')::jsonb, infrastructure),
    updated_at     = now()
WHERE id = sqlc.arg('id')
RETURNING *;
