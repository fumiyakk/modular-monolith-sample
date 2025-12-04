package get_user

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/entity"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/repository"
	"github.com/google/uuid"
)

type Usecase interface {
	GetUser(context.Context, uuid.UUID) (*entity.User, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) Usecase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (u *usecase) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := u.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
