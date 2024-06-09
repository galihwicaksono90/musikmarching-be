package postgres

import (
	"context"
	"fmt"
	"os"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	"github.com/jackc/pgx/v5"
)

func DB(ctx *context.Context) *db.Queries {
	dbUrl := "postgres://tony:password@localhost:5432/swaranada-be"

	conn, err := pgx.Connect(*ctx, dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(*ctx)

	queries := db.New(conn)
	return queries
}
