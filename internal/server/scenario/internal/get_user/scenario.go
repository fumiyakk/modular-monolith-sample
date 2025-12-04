package get_user

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/entity"
	"github.com/google/uuid"
)

type Scenario interface {
	GetUser(context.Context, uuid.UUID) (*entity.User, error)
}

type scenario struct {
	uow         unit_of_work.UnitOfWork
	userUsecase user.GetUserUsecase
}

func New(
	uow unit_of_work.UnitOfWork,
	userUsecase user.GetUserUsecase,
) Scenario {
	return &scenario{
		uow:         uow,
		userUsecase: userUsecase,
	}
}

func (s *scenario) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user *entity.User

	err := s.uow.WithinTransaction(ctx, func(ctx context.Context) error {
		u, err := s.userUsecase.GetUser(ctx, id)
		if err != nil {
			return err
		}
		user = u

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
