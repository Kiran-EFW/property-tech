-- name: CreateLead :one
INSERT INTO leads (
    id, investor_id, project_id, agent_id, source,
    status, phone, name, email, budget, notes,
    follow_up_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.narg('investor_id'),
    sqlc.arg('project_id'),
    sqlc.narg('agent_id'),
    sqlc.arg('source'),
    'new',
    sqlc.arg('phone'),
    sqlc.arg('name'),
    sqlc.narg('email'),
    sqlc.narg('budget'),
    sqlc.narg('notes'),
    sqlc.narg('follow_up_at'),
    now(),
    now()
) RETURNING *;

-- name: GetLeadByID :one
SELECT * FROM leads
WHERE id = sqlc.arg('id');

-- name: ListLeads :many
SELECT * FROM leads
WHERE (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
  AND (sqlc.narg('agent_id')::uuid IS NULL OR agent_id = sqlc.narg('agent_id'))
  AND (sqlc.narg('project_id')::uuid IS NULL OR project_id = sqlc.narg('project_id'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListLeadsByAgent :many
SELECT * FROM leads
WHERE agent_id = sqlc.arg('agent_id')
  AND (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateLeadStatus :one
UPDATE leads SET
    status     = sqlc.arg('status'),
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: AssignLeadToAgent :one
UPDATE leads SET
    agent_id   = sqlc.arg('agent_id'),
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpdateLeadFollowUp :exec
UPDATE leads SET
    follow_up_at = sqlc.arg('follow_up_at'),
    updated_at   = now()
WHERE id = sqlc.arg('id');

-- name: CountLeadsByStatus :many
SELECT status, count(*) AS count
FROM leads
WHERE (sqlc.narg('agent_id')::uuid IS NULL OR agent_id = sqlc.narg('agent_id'))
GROUP BY status;

-- name: FindLeadByPhone :one
SELECT * FROM leads
WHERE phone = sqlc.arg('phone')
ORDER BY created_at DESC
LIMIT 1;

-- name: ListOverdueFollowUps :many
SELECT * FROM leads
WHERE follow_up_at < now()
  AND status NOT IN ('converted', 'lost')
  AND (sqlc.narg('agent_id')::uuid IS NULL OR agent_id = sqlc.narg('agent_id'))
ORDER BY follow_up_at ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
