-- name: CreateSnippet :one 
INSERT INTO snippets(content,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSnippetByID :one
SELECT * FROM snippets 
WHERE id = $1
LIMIT 1;

-- name: ListSnippets :many
SELECT * FROM snippets
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListSnippetsByUser :many
SELECT * FROM snippets 
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateSnippetContent :exec
UPDATE snippets
SET content = $1, updated_at = now()
WHERE id = $2;

-- name: DeleteSnippet :exec
DELETE FROM snippets
WHERE id = $1;
