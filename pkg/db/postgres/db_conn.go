package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func New(dbSource string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), dbSource)
}
