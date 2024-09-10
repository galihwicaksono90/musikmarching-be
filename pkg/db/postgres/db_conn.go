package postgres

import (
	"context"
	"fmt"
	"os"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DB(ctx *context.Context) *db.Queries {
	dbUrl := "postgres://admin:root@localhost:5432/musikmarching-db"

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// defer conn.af(*ctx)
	// config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// 	fmt.Println("Connected to database")
	// 	return nil
	// }

	pool, err := pgxpool.NewWithConfig(*ctx, config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	queries := db.New(pool)
	return queries
}
