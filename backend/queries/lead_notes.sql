-- name: CreateLeadNote :one
INSERT INTO lead_notes (
    id, lead_id, author_id, content, created_at
) VALUES (
    gen_random_uuid(),
    sqlc.arg('lead_id'),
    sqlc.arg('author_id'),
    sqlc.arg('content'),
    now()
) RETURNING *;

-- name: ListNotesByLead :many
SELECT * FROM lead_notes
WHERE lead_id = sqlc.arg('lead_id')
ORDER BY created_at DESC;
