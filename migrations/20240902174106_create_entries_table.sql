-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entries (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  task_id INTEGER NOT NULL,
  FOREIGN KEY (task_id) REFERENCES tasks (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS entries;
-- +goose StatementEnd
