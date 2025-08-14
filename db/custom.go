package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, context.Context) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STR"))
	if err != nil {
		log.Fatal("Something went wrong")
	}

	return conn, ctx
}
