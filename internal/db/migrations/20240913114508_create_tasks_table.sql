-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  project_id INTEGER NOT NULL,
  FOREIGN KEY (project_id) REFERENCES projects (id),
  UNIQUE(project_id, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
