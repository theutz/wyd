package queries

import (
	_ "github.com/sqlc-dev/sqlc" // needed for go generate
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
