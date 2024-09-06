-- name: ProjectsCount :one
SELECT COUNT(*)
FROM projects;

-- name: AddProject :one
INSERT INTO projects (name)
VALUES (?)
RETURNING *;
