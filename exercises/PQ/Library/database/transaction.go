package database

import (
	"context"
	"excercise-library/ent"
	"excercise-library/shared"
)

type EntClientTx struct {
	Tx *ent.Tx
}

type contextKey struct {
	key string
}

var Key = contextKey{key: "entTX"}

type Transaction interface {
	StartTx(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func CommitOrRollback(ctx context.Context, t Transaction, err error) error {
	if err != nil {
		// error occurred , try rollback
		if err2 := t.Rollback(ctx); err2 != nil {
			// rollback failed
			return shared.NewApiError(err2.Error(), shared.Internal, "CommitOrRollback", shared.TxMiddlewareLevel, err)
		}
		// Rollback succeed
		return err
	}
	// no error, proceed to commit
	err = t.Commit(ctx)
	// happy path, commit succeed
	if err != nil {
		// commit fail, try rollback
		if err2 := t.Rollback(ctx); err2 != nil {
			// rollback failed
			return shared.NewApiError(err2.Error(), shared.Internal, "CommitOrRollback", shared.TxMiddlewareLevel, err)
		}
	}
	return err
}

type RepositoryTx interface {
	Transaction
	GetClient(ctx context.Context) *ent.Client
}

type RepositoryTxImpl struct {
	Client *ent.Client
}

func (r *RepositoryTxImpl) GetClient(ctx context.Context) *ent.Client {
	if client := ctx.Value(Key); client != nil {
		return client.(EntClientTx).Tx.Client()
	}
	return r.Client
}

func (r RepositoryTxImpl) getTx(ctx context.Context) (*ent.Tx, error) {
	tx := ctx.Value(Key)
	if tx != nil {
		return tx.(EntClientTx).Tx, nil
	}
	return nil, shared.ErrDatabaseTx
}

func (r RepositoryTxImpl) StartTx(ctx context.Context) (context.Context, error) {
	tx, err := r.GetClient(ctx).Tx(ctx)
	if err == nil {
		ctx = context.WithValue(ctx, Key, EntClientTx{
			Tx: tx,
		})
	}
	return ctx, err
}

func (r RepositoryTxImpl) Commit(ctx context.Context) error {
	tx, err := r.getTx(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r RepositoryTxImpl) Rollback(ctx context.Context) error {
	tx, err := r.getTx(ctx)
	if err != nil {
		return err
	}
	return tx.Rollback()
}
