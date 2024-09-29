package persistence

import (
	"sync"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *pgx.Conn
	mu sync.Mutex
	*Queries
}

// func NewStoree(db *pgxpool.Conn) Store {
// 	return &SQLStore{
// 		db:      db,
// 		Queries: New(db),
// 	}
// }

func NewStore(db *pgx.Conn) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	store.mu.Lock()
// 	defer store.mu.Unlock()
// 	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}
//
// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}
//
// 	return tx.Commit(ctx)
// }
