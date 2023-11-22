package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/ToroEtele/go-graphql-api/internal/database"
)

type Config struct {
	Db   *database.Queries
	Conn *sql.DB
}

func InitDb() (Config, error) {
	var err error
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return Config{
			Db:   nil,
			Conn: nil,
		}, errors.New("DB_URL environment variable is not set")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Println(err)
		return Config{
			Db:   nil,
			Conn: nil,
		}, errors.New("can't connect to database")
	}

	db := database.New(conn)

	return Config{
		Db:   db,
		Conn: conn,
	}, nil
}
