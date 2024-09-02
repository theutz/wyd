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

-- name: CreateTask :one
INSERT INTO tasks (name, project_id)
VALUES (?, ?)
RETURNING *;

-- name: ListTasks :many
SELECT t.name, p.name AS project_name
FROM tasks AS t
INNER JOIN projects AS p
ON p.id = t.project_id;

-- name: ListEntries :many
SELECT e.name, t.name AS task_name
FROM entries AS e
INNER JOIN tasks AS t
ON t.id = e.task_id;
