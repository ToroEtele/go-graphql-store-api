package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ToroEtele/go-graphql-api/cmd/app/resolvers"
	"github.com/ToroEtele/go-graphql-api/config"
	"github.com/ToroEtele/go-graphql-api/graph"
	"github.com/ToroEtele/go-graphql-api/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := config.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.AuthMiddleware)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		DB: db,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
