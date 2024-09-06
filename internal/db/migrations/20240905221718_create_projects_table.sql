-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS projects;
-- +goose StatementEnd
