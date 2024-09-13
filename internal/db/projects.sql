-- name: AddProject :one
INSERT INTO projects (name, client_id)
VALUES (?, ?)
RETURNING *;

-- name: ListProjects :many
SELECT *
FROM projects;

-- name: DeleteProject :one
DELETE FROM projects
WHERE id = ?
RETURNING *;

