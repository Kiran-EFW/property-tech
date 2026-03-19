-- name: CreateUnit :one
INSERT INTO project_units (
    id, project_id, unit_number, floor, unit_type,
    carpet_area, price, status, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('project_id'),
    sqlc.arg('unit_number'),
    sqlc.narg('floor'),
    sqlc.arg('unit_type'),
    sqlc.arg('carpet_area'),
    sqlc.arg('price'),
    COALESCE(sqlc.narg('status')::varchar, 'available'),
    now(),
    now()
) RETURNING *;

-- name: GetUnitByID :one
SELECT * FROM project_units
WHERE id = sqlc.arg('id');

-- name: ListUnitsByProject :many
SELECT * FROM project_units
WHERE project_id = sqlc.arg('project_id')
  AND (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
ORDER BY floor, unit_number;

-- name: UpdateUnit :one
UPDATE project_units SET
    unit_number = COALESCE(sqlc.narg('unit_number')::varchar, unit_number),
    floor       = COALESCE(sqlc.narg('floor')::int, floor),
    unit_type   = COALESCE(sqlc.narg('unit_type')::varchar, unit_type),
    carpet_area = COALESCE(sqlc.narg('carpet_area')::float8, carpet_area),
    price       = COALESCE(sqlc.narg('price')::float8, price),
    status      = COALESCE(sqlc.narg('status')::varchar, status),
    updated_at  = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpdateUnitStatus :exec
UPDATE project_units SET
    status     = sqlc.arg('status'),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: CountAvailableUnits :one
SELECT count(*) FROM project_units
WHERE project_id = sqlc.arg('project_id')
  AND status = 'available';
