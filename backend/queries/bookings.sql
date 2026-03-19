-- name: CreateBooking :one
INSERT INTO bookings (
    id, project_id, unit_id, investor_id, agent_id,
    booking_amount, agreement_value, status,
    booked_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('project_id'),
    sqlc.narg('unit_id'),
    sqlc.arg('investor_id'),
    sqlc.narg('agent_id'),
    sqlc.arg('booking_amount'),
    sqlc.arg('agreement_value'),
    'pending',
    now(),
    now(),
    now()
) RETURNING *;

-- name: GetBookingByID :one
SELECT * FROM bookings
WHERE id = sqlc.arg('id');

-- name: ListBookings :many
SELECT * FROM bookings
WHERE (sqlc.narg('status')::varchar IS NULL OR status = sqlc.narg('status'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListBookingsByAgent :many
SELECT * FROM bookings
WHERE agent_id = sqlc.arg('agent_id')
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListBookingsByInvestor :many
SELECT * FROM bookings
WHERE investor_id = sqlc.arg('investor_id')
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateBookingStatus :one
UPDATE bookings SET
    status     = sqlc.arg('status'),
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: CountBookingsByStatus :many
SELECT status, count(*) AS count
FROM bookings
GROUP BY status;
