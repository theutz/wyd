-- name: CreateClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;
