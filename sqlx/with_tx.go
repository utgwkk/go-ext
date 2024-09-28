package sqlxsnippet

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func WithTx(ctx context.Context, db *sqlx.DB, f func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to BeginTxx: %w", err)
	}
	defer tx.Rollback()

	if err := f(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to Commit: %w", err)
	}

	return nil
}

func WithTx2[T any](ctx context.Context, db *sqlx.DB, f func(tx *sqlx.Tx) (T, error)) (T, error) {
	var t T
	f2 := func(tx *sqlx.Tx) error {
		_t, err := f(tx)
		if err != nil {
			return err
		}
		t = _t
		return nil
	}

	if err := WithTx(ctx, db, f2); err != nil {
		return t, err
	}
	return t, nil
}
