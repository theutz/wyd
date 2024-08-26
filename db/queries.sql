-- name: CreateClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: ListClients :many
SELECT * from clients;
