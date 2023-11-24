package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	model "github.com/ToroEtele/go-graphql-api/cmd/app/domain"
	"github.com/ToroEtele/go-graphql-api/middleware"
)

func StockModifier(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	currentUser := ctx.Value(middleware.ContextKey("user"))

	if val, ok := currentUser.(*model.CurrentUser); ok {
		if val.IsAdmin {
			return next(ctx)
		}

	}
	stock, _ := next(ctx)
	if val, ok := stock.(int); ok && val > 1 {
		return 2, nil
	}
	return stock, nil

}
