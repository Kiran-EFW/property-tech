-- name: CreateCommission :one
INSERT INTO commissions (
    id, booking_id, agent_id, amount, tds, net_amount,
    status, paid_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('booking_id'),
    sqlc.arg('agent_id'),
    sqlc.arg('amount'),
    sqlc.arg('tds'),
    sqlc.arg('net_amount'),
    'pending',
    NULL,
    now(),
    now()
) RETURNING *;

-- name: GetCommissionByID :one
SELECT * FROM commissions
WHERE id = sqlc.arg('id');

-- name: ListCommissions :many
SELECT * FROM commissions
WHERE (sqlc.narg('agent_id')::uuid IS NULL OR agent_id = sqlc.narg('agent_id'))
  AND (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListCommissionsByAgent :many
SELECT * FROM commissions
WHERE agent_id = sqlc.arg('agent_id')
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateCommissionStatus :one
UPDATE commissions SET
    status     = sqlc.arg('status'),
    paid_at    = CASE WHEN sqlc.arg('status')::varchar = 'paid' THEN now() ELSE paid_at END,
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: GetCommissionSummary :one
SELECT
    COALESCE(SUM(amount), 0)::float8     AS total_amount,
    COALESCE(SUM(tds), 0)::float8        AS total_tds,
    COALESCE(SUM(net_amount), 0)::float8 AS total_net_amount,
    count(*)                              AS total_count
FROM commissions
WHERE agent_id = sqlc.arg('agent_id')
  AND (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'));
