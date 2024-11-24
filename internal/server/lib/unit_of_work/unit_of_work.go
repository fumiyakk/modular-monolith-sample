package unit_of_work

import (
	"context"
	"fmt"
	"sync"
)

type ctxKey struct{}

type UnitOfWork interface {
	WithinTransaction(ctx context.Context, f func(ctx context.Context) error) error
}

type Transaction interface {
	Commit() error
	Rollback() error
	GetChanges() map[string]interface{}
	SetChange(key string, value interface{})
	IsActive() bool
	ApplyChanges(apply func(changes map[string]interface{}) error) error
}

type unitOfWork struct {
	mu sync.RWMutex
}

type transaction struct {
	mu        sync.RWMutex
	committed bool
	rollback  bool
	changes   map[string]interface{}
}

func NewUnitOfWork() UnitOfWork {
	return &unitOfWork{}
}

func GetTx(ctx context.Context) (Transaction, bool) {
	tx, ok := ctx.Value(ctxKey{}).(Transaction)
	return tx, ok
}

func contextWithTx(ctx context.Context, tx Transaction) context.Context {
	return context.WithValue(ctx, ctxKey{}, tx)
}

func (uow *unitOfWork) WithinTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	if _, ok := GetTx(ctx); ok {
		return f(ctx)
	}

	tx := &transaction{
		changes: make(map[string]interface{}),
	}

	txCtx := contextWithTx(ctx, tx)

	if err := f(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (tx *transaction) GetChanges() map[string]interface{} {
	tx.mu.RLock()
	defer tx.mu.RUnlock()
	return tx.changes
}

func (tx *transaction) SetChange(key string, value interface{}) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.changes[key] = value
}

func (tx *transaction) IsActive() bool {
	tx.mu.RLock()
	defer tx.mu.RUnlock()
	return !tx.committed && !tx.rollback
}

func (tx *transaction) Commit() error {
	tx.mu.Lock()
	defer tx.mu.Unlock()

	if tx.committed {
		return fmt.Errorf("transaction already committed")
	}
	if tx.rollback {
		return fmt.Errorf("transaction already rolled back")
	}

	tx.committed = true
	return nil
}

func (tx *transaction) Rollback() error {
	tx.mu.Lock()
	defer tx.mu.Unlock()

	if tx.committed {
		return fmt.Errorf("transaction already committed")
	}
	if tx.rollback {
		return fmt.Errorf("transaction already rolled back")
	}

	tx.rollback = true
	return nil
}

func (tx *transaction) ApplyChanges(apply func(changes map[string]interface{}) error) error {
	tx.mu.RLock()
	defer tx.mu.RUnlock()

	if tx.committed {
		return fmt.Errorf("transaction already committed")
	}
	if tx.rollback {
		return fmt.Errorf("transaction already rolled back")
	}

	return apply(tx.changes)
}
