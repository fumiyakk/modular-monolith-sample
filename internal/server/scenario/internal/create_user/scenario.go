package create_user

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user"
	"github.com/google/uuid"
)

type Scenario interface {
	CreateUser(context.Context, string) (uuid.UUID, uuid.UUID, error)
}

type scenario struct {
	uow             unit_of_work.UnitOfWork
	userUsecase     user.CreateUserUsecase
	contractUsecase contract.CreateContractUsecase
}

func New(
	uow unit_of_work.UnitOfWork,
	userUsecase user.CreateUserUsecase,
	contractUsecase contract.CreateContractUsecase,
) Scenario {
	return &scenario{
		uow:             uow,
		userUsecase:     userUsecase,
		contractUsecase: contractUsecase,
	}
}

func (s *scenario) CreateUser(ctx context.Context, name string) (uuid.UUID, uuid.UUID, error) {
	var userID, contractID uuid.UUID

	err := s.uow.WithinTransaction(ctx, func(ctx context.Context) error {
		id, err := s.userUsecase.CreateUser(ctx, name)
		if err != nil {
			return err
		}
		userID = id

		contractID, err = s.contractUsecase.CreateContract(ctx, userID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return userID, contractID, nil
}
