package user

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/usecase/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user/internal/usecase/get_user"
)

type CreateUserUsecase = create_user.Usecase
type GetUserUsecase = get_user.Usecase
