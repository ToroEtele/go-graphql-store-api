package resolvers

import "github.com/ToroEtele/go-graphql-api/internal/database"

type Resolver struct {
	DB *database.Queries
}
