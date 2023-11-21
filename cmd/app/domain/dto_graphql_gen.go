// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	model "github.com/ToroEtele/go-graphql-api/cmd/app/domain/dao"
)

type CreateProductInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Stock       int     `json:"stock"`
}

type CurrentUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

type OrdersFilter struct {
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	IsPaid  *bool   `json:"isPaid,omitempty"`
	Country *string `json:"country,omitempty"`
	Address *string `json:"address,omitempty"`
	UserID  *string `json:"userId,omitempty"`
}

type OrdersResponse struct {
	Edges      []*model.Order `json:"edges"`
	TotalCount int            `json:"totalCount"`
}

type Pagination struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateProductInput struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Description *string  `json:"description,omitempty"`
	Image       *string  `json:"image,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}

type UsersFilter struct {
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	IsAdmin *bool   `json:"isAdmin,omitempty"`
}

type UsersResponse struct {
	Edges      []*model.User `json:"edges"`
	TotalCount int           `json:"totalCount"`
}
