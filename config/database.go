package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/ToroEtele/go-graphql-api/internal/database"
)

func InitDb() (*database.Queries, error) {
	var err error
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL environment variable is not set")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can't connect to database")
	}

	db := database.New(conn)

	return db, nil
}
