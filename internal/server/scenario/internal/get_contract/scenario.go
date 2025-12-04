package get_contract

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/entity"
	"github.com/google/uuid"
)

type Scenario interface {
	GetContract(context.Context, uuid.UUID) (*entity.Contract, error)
}

type scenario struct {
	uow             unit_of_work.UnitOfWork
	contractUsecase contract.GetContractUsecase
}

func New(
	uow unit_of_work.UnitOfWork,
	contractUsecase contract.GetContractUsecase,
) Scenario {
	return &scenario{
		uow:             uow,
		contractUsecase: contractUsecase,
	}
}

func (s *scenario) GetContract(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	var contract *entity.Contract

	err := s.uow.WithinTransaction(ctx, func(ctx context.Context) error {
		u, err := s.contractUsecase.GetContract(ctx, id)
		if err != nil {
			return err
		}
		contract = u

		return nil
	})

	if err != nil {
		return nil, err
	}

	return contract, nil
}
