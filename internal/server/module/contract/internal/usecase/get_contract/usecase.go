package get_contract

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/entity"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/repository"
	"github.com/google/uuid"
)

type Usecase interface {
	GetContract(context.Context, uuid.UUID) (*entity.Contract, error)
}

type usecase struct {
	contractRepo repository.ContractRepository
}

func New(contractRepo repository.ContractRepository) Usecase {
	return &usecase{
		contractRepo: contractRepo,
	}
}

func (u *usecase) GetContract(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	contract, err := u.contractRepo.GetContract(ctx, id)
	if err != nil {
		return nil, err
	}
	return contract, nil
}
