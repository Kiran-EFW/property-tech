-- name: CreateBuilder :one
INSERT INTO builders (
    id, name, slug, rera_number, pan, gst,
    track_record, contact_phone, contact_email,
    logo_url, is_verified, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('name'),
    sqlc.arg('slug'),
    sqlc.arg('rera_number'),
    sqlc.narg('pan'),
    sqlc.narg('gst'),
    sqlc.narg('track_record'),
    sqlc.narg('contact_phone'),
    sqlc.narg('contact_email'),
    sqlc.narg('logo_url'),
    COALESCE(sqlc.narg('is_verified')::boolean, false),
    now(),
    now()
) RETURNING *;

-- name: GetBuilderByID :one
SELECT * FROM builders
WHERE id = sqlc.arg('id');

-- name: GetBuilderBySlug :one
SELECT * FROM builders
WHERE slug = sqlc.arg('slug');

-- name: ListBuilders :many
SELECT * FROM builders
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateBuilder :one
UPDATE builders SET
    name          = COALESCE(sqlc.narg('name')::varchar, name),
    slug          = COALESCE(sqlc.narg('slug')::varchar, slug),
    rera_number   = COALESCE(sqlc.narg('rera_number')::varchar, rera_number),
    pan           = COALESCE(sqlc.narg('pan')::varchar, pan),
    gst           = COALESCE(sqlc.narg('gst')::varchar, gst),
    track_record  = COALESCE(sqlc.narg('track_record')::text, track_record),
    contact_phone = COALESCE(sqlc.narg('contact_phone')::varchar, contact_phone),
    contact_email = COALESCE(sqlc.narg('contact_email')::varchar, contact_email),
    logo_url      = COALESCE(sqlc.narg('logo_url')::varchar, logo_url),
    is_verified   = COALESCE(sqlc.narg('is_verified')::boolean, is_verified),
    updated_at    = now()
WHERE id = sqlc.arg('id')
RETURNING *;
