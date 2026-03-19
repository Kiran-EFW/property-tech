-- name: CreateSiteVisit :one
INSERT INTO site_visits (
    id, lead_id, agent_id, project_id,
    feedback, photos, duration, rating,
    visited_at, created_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('lead_id'),
    sqlc.arg('agent_id'),
    sqlc.arg('project_id'),
    sqlc.narg('feedback'),
    COALESCE(sqlc.narg('photos')::jsonb, '[]'::jsonb),
    sqlc.narg('duration'),
    sqlc.narg('rating'),
    sqlc.arg('visited_at'),
    now()
) RETURNING *;

-- name: GetSiteVisitByID :one
SELECT * FROM site_visits
WHERE id = sqlc.arg('id');

-- name: ListSiteVisits :many
SELECT * FROM site_visits
WHERE (sqlc.narg('agent_id')::uuid IS NULL OR agent_id = sqlc.narg('agent_id'))
ORDER BY visited_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListSiteVisitsByLead :many
SELECT * FROM site_visits
WHERE lead_id = sqlc.arg('lead_id')
ORDER BY visited_at DESC;

-- name: UpdateSiteVisitFeedback :one
UPDATE site_visits SET
    feedback = sqlc.narg('feedback'),
    photos   = COALESCE(sqlc.narg('photos')::jsonb, photos),
    duration = COALESCE(sqlc.narg('duration')::int, duration),
    rating   = COALESCE(sqlc.narg('rating')::int, rating)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: CountVisitsByAgent :one
SELECT count(*) FROM site_visits
WHERE agent_id = sqlc.arg('agent_id');
