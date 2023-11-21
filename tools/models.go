package tools

import (
	model "github.com/ToroEtele/go-graphql-api/cmd/app/domain/dao"
	"github.com/ToroEtele/go-graphql-api/internal/database"
)

func DatabaseProductsToProducts(products []database.Product) []*model.Product {
	databaseProducts := []*model.Product{}
	for _, product := range products {
		databaseProducts = append(databaseProducts, DatabaseProductToProduct(product))
	}
	return databaseProducts
}

func DatabaseProductToProduct(product database.Product) *model.Product {
	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       int(product.Stock),
		Image:       product.Image,
		Description: product.Description,
	}
}
