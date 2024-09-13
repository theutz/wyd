-- name: AddClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: ListClients :many
SELECT *
FROM clients;

-- name: DeleteClient :one
DELETE FROM clients
WHERE id = ?
RETURNING *;

-- name: DeleteClientByName :one
DELETE FROM clients
WHERE name LIKE ?
RETURNING *;

