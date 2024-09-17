-- name: All :many
SELECT *
FROM clients;

-- name: Create :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;
