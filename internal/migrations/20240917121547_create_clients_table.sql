-- +goose Up
-- +goose StatementBegin
CREATE IF NOT EXISTS clients (
  id INT PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP IF EXISTS clients;
-- +goose StatementEnd
