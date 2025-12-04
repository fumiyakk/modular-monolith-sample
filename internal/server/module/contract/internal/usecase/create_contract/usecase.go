package create_contract

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/entity"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/repository"
	"github.com/google/uuid"
)

type Usecase interface {
	CreateContract(context.Context, uuid.UUID) (uuid.UUID, error)
}

type usecase struct {
	contractRepo repository.ContractRepository
}

func New(contractRepo repository.ContractRepository) Usecase {
	return &usecase{
		contractRepo: contractRepo,
	}
}

func (u *usecase) CreateContract(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	contract := &entity.Contract{
		ID:     uuid.New(),
		UserID: userID,
		Status: entity.ContractStatusActive,
	}

	err := u.contractRepo.CreateContract(ctx, contract)
	if err != nil {
		return uuid.Nil, err
	}
	return contract.ID, nil
}
