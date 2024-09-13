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

-- name: AddTask :one
INSERT INTO tasks (name, project_id)
VALUES (?, ?)
RETURNING *;

-- name: ListTasks :many
SELECT *
FROM tasks;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = ?
RETURNING *;
