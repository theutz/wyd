-- name: AddClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: ListClients :many
SELECT *
FROM clients;

-- name: DeleteClient :one
DELETE FROM clients
WHERE name LIKE ?
RETURNING *;

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
