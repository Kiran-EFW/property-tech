-- name: CreateProject :one
INSERT INTO projects (
    id, name, slug, rera_number, builder_id, description,
    carpet_area_min, carpet_area_max, price_min, price_max,
    location, address, city, state, pincode,
    status, amenities, media, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('name'),
    sqlc.arg('slug'),
    sqlc.arg('rera_number'),
    sqlc.arg('builder_id'),
    sqlc.narg('description'),
    sqlc.narg('carpet_area_min'),
    sqlc.narg('carpet_area_max'),
    sqlc.narg('price_min'),
    sqlc.narg('price_max'),
    ST_MakePoint(sqlc.arg('lng')::float8, sqlc.arg('lat')::float8)::geography,
    sqlc.narg('address'),
    sqlc.narg('city'),
    sqlc.narg('state'),
    sqlc.narg('pincode'),
    COALESCE(sqlc.narg('status')::varchar, 'draft'),
    COALESCE(sqlc.narg('amenities')::jsonb, '[]'::jsonb),
    COALESCE(sqlc.narg('media')::jsonb, '[]'::jsonb),
    now(),
    now()
) RETURNING *;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = sqlc.arg('id');

-- name: GetProjectBySlug :one
SELECT * FROM projects
WHERE slug = sqlc.arg('slug');

-- name: ListProjects :many
SELECT * FROM projects
WHERE (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
  AND (sqlc.narg('city')::varchar IS NULL OR city = sqlc.narg('city'))
  AND (sqlc.narg('price_min_filter')::float8 IS NULL OR price_min >= sqlc.narg('price_min_filter'))
  AND (sqlc.narg('price_max_filter')::float8 IS NULL OR price_max <= sqlc.narg('price_max_filter'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateProject :one
UPDATE projects SET
    name           = COALESCE(sqlc.narg('name')::varchar, name),
    slug           = COALESCE(sqlc.narg('slug')::varchar, slug),
    rera_number    = COALESCE(sqlc.narg('rera_number')::varchar, rera_number),
    description    = COALESCE(sqlc.narg('description')::text, description),
    carpet_area_min = COALESCE(sqlc.narg('carpet_area_min')::float8, carpet_area_min),
    carpet_area_max = COALESCE(sqlc.narg('carpet_area_max')::float8, carpet_area_max),
    price_min      = COALESCE(sqlc.narg('price_min')::float8, price_min),
    price_max      = COALESCE(sqlc.narg('price_max')::float8, price_max),
    address        = COALESCE(sqlc.narg('address')::varchar, address),
    city           = COALESCE(sqlc.narg('city')::varchar, city),
    state          = COALESCE(sqlc.narg('state')::varchar, state),
    pincode        = COALESCE(sqlc.narg('pincode')::varchar, pincode),
    status         = COALESCE(sqlc.narg('status')::varchar, status),
    amenities      = COALESCE(sqlc.narg('amenities')::jsonb, amenities),
    media          = COALESCE(sqlc.narg('media')::jsonb, media),
    updated_at     = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: CountProjects :one
SELECT count(*) FROM projects
WHERE (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'));

-- name: ListProjectsByBuilder :many
SELECT * FROM projects
WHERE builder_id = sqlc.arg('builder_id')
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: SearchProjectsByLocation :many
SELECT * FROM projects
WHERE ST_DWithin(
    location,
    ST_MakePoint(sqlc.arg('lng'), sqlc.arg('lat'))::geography,
    sqlc.arg('radius_meters')
)
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
