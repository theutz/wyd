-- name: ProjectsCount :one
SELECT COUNT(*)
FROM projects;

-- name: AddProject :one
INSERT INTO projects (name, client_id)
VALUES (?, ?)
RETURNING *;

-- name: AddClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;
