package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type userRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*entity.User
}

func New() UserRepository {
	return &userRepository{
		users: make(map[uuid.UUID]*entity.User),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	tx, ok := unit_of_work.GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction in context")
	}

	if !tx.IsActive() {
		return fmt.Errorf("transaction is no longer active")
	}

	// Check if user already exists
	r.mu.RLock()
	_, exists := r.users[user.ID]
	r.mu.RUnlock()
	if exists {
		return fmt.Errorf("user already exists: %v", user.ID)
	}

	// Store in transaction
	userCopy := *user
	tx.SetChange(fmt.Sprintf("user:%s", user.ID), &userCopy)

	// Apply changes if this is the outermost transaction
	return tx.ApplyChanges(func(changes map[string]interface{}) error {
		r.mu.Lock()
		defer r.mu.Unlock()

		for _, value := range changes {
			if user, ok := value.(*entity.User); ok {
				r.users[user.ID] = user
			}
		}
		return nil
	})
}

func (r *userRepository) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	tx, ok := unit_of_work.GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("no transaction in context")
	}

	if !tx.IsActive() {
		return nil, fmt.Errorf("transaction is no longer active")
	}

	// Check transaction changes first
	changes := tx.GetChanges()
	if userInterface, exists := changes[fmt.Sprintf("user:%s", id)]; exists {
		user := userInterface.(*entity.User)
		userCopy := *user
		return &userCopy, nil
	}

	// Check main storage
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found: %v", id)
	}

	userCopy := *user
	return &userCopy, nil
}
