-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clients (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS tmp_projects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  client_id INTEGER NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients (id)
);

DROP TABLE projects;

ALTER TABLE tmp_projects RENAME TO projects;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
