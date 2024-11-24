package user

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/repository"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/usecase/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/usecase/get_user"
)

type Module interface {
	CreateUserUsecase
	GetUserUsecase
}

func New() Module {
	repo := repository.New()

	return &struct {
		CreateUserUsecase
		GetUserUsecase
	}{
		create_user.New(repo),
		get_user.New(repo),
	}
}
