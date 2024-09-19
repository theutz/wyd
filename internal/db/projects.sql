-- name: All :many
SELECT *
FROM projects;

-- name: Create :one
INSERT INTO projects (name, client_id)
VALUES (?, ?)
RETURNING *;

-- name: DeleteByName :one
DELETE FROM projects
WHERE name LIKE ?
RETURNING *;
