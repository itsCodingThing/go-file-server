package db

import (
	"context"
	"fmt"
	"go-file-server/db/sqlc"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	Connection *pgx.Conn
	Queries    *sqlc.Queries
)

func Connect() (c *pgx.Conn) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Connection = conn
	Queries = sqlc.New(conn)

	fmt.Println("database is contected 🚀")

	return conn
}
