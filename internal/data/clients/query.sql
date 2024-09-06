-- name: AddClient :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;
