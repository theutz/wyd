-- name: CreateClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: ListClients :many
SELECT * from clients;

-- name: CreateProject :one
INSERT INTO projects (name, client_id)
VALUES (?, ?)
RETURNING *;

-- name: ListProjects :many
SELECT * from projects;
