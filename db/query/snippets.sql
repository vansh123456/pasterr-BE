-- name: CreateSnippet :one 
INSERT INTO snippets(content,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSnippetByID :one
SELECT * FROM snippets 
WHERE id = $1;

-- name: ListSnippets :many
SELECT * FROM snippets
ORDER BY created_at DESC;


-- name: ListSnippetsByUserID :many
SELECT id, content, user_id, created_at, updated_at
FROM snippets
WHERE user_id = $1;


-- name: UpdateSnippetContent :one
UPDATE snippets
SET content = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteSnippet :exec
DELETE FROM snippets
WHERE id = $1;