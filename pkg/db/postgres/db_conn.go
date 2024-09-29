package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func New(ctx context.Context, dbSource string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dbSource)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
