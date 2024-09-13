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
