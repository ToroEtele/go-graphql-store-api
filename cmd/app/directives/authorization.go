package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	model "github.com/ToroEtele/go-graphql-api/cmd/app/domain"
	"github.com/ToroEtele/go-graphql-api/middleware"
)

func Authorization(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	currentUser := ctx.Value(middleware.ContextKey("user"))

	if val, ok := currentUser.(*model.CurrentUser); ok {
		if val.IsAdmin {
			return next(ctx)
		}
	}

	return nil, errors.New("Unauthorized")

}
