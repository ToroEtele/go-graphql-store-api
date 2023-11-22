package resolvers

import (
	"database/sql"

	"github.com/ToroEtele/go-graphql-api/internal/database"
)

type Resolver struct {
	DB   *database.Queries
	Conn *sql.DB
}
