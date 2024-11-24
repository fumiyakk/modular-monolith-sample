package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/entity"

	"github.com/google/uuid"
)

type ContractRepository interface {
	CreateContract(ctx context.Context, contract *entity.Contract) error
	GetContract(ctx context.Context, id uuid.UUID) (*entity.Contract, error)
}

type contractRepository struct {
	mu        sync.RWMutex
	contracts map[uuid.UUID]*entity.Contract
}

func New() ContractRepository {
	return &contractRepository{
		contracts: make(map[uuid.UUID]*entity.Contract),
	}
}

func (r *contractRepository) CreateContract(ctx context.Context, contract *entity.Contract) error {
	tx, ok := unit_of_work.GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction in context")
	}

	if !tx.IsActive() {
		return fmt.Errorf("transaction is no longer active")
	}

	// Check if contract already exists
	r.mu.RLock()
	_, exists := r.contracts[contract.ID]
	r.mu.RUnlock()
	if exists {
		return fmt.Errorf("contract already exists: %v", contract.ID)
	}

	// Store in transaction
	contractCopy := *contract
	tx.SetChange(fmt.Sprintf("contract:%s", contract.ID), &contractCopy)

	// Apply changes if this is the outermost transaction
	return tx.ApplyChanges(func(changes map[string]interface{}) error {
		r.mu.Lock()
		defer r.mu.Unlock()

		for _, value := range changes {
			if contract, ok := value.(*entity.Contract); ok {
				r.contracts[contract.ID] = contract
			}
		}
		return nil
	})
}

func (r *contractRepository) GetContract(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	tx, ok := unit_of_work.GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("no transaction in context")
	}

	if !tx.IsActive() {
		return nil, fmt.Errorf("transaction is no longer active")
	}

	// Check transaction changes first
	changes := tx.GetChanges()
	if contractInterface, exists := changes[fmt.Sprintf("contract:%s", id)]; exists {
		contract := contractInterface.(*entity.Contract)
		contractCopy := *contract
		return &contractCopy, nil
	}

	// Check main storage
	r.mu.RLock()
	defer r.mu.RUnlock()

	contract, exists := r.contracts[id]
	if !exists {
		return nil, fmt.Errorf("contract not found: %v", id)
	}

	contractCopy := *contract
	return &contractCopy, nil
}
