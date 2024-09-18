-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  client_id INTEGER NOT NULL,
  UNIQUE (name, client_id),
  FOREIGN KEY (client_id) REFERENCES clients (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS projects;
-- +goose StatementEnd
