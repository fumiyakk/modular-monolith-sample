package create_user

import (
	"context"

	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/entity"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/repository"
	"github.com/google/uuid"
)

type Usecase interface {
	CreateUser(context.Context, string) (uuid.UUID, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) Usecase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (u *usecase) CreateUser(ctx context.Context, name string) (uuid.UUID, error) {
	user := &entity.User{
		ID:   uuid.New(),
		Name: name,
	}

	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}
