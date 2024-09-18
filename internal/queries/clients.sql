-- name: All :many
SELECT *
FROM clients;

-- name: Create :one
INSERT INTO clients (name)
VALUES (?)
RETURNING *;

-- name: DeleteByName :one
DELETE FROM clients
WHERE name LIKE ?
RETURNING *;
