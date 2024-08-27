-- name: CreateClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: ListClients :many
SELECT * from clients;

-- name: GetClientByName :one
SELECT *
FROM clients
WHERE name LIKE ?;

-- name: CreateProject :one
INSERT INTO projects (name, client_id)
VALUES (?, ?)
RETURNING *;

-- name: ListProjects :many
SELECT p.*, c.name AS client_name
FROM projects AS p
INNER JOIN clients AS c
ON c.id = p.client_id;
