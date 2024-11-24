package sample

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/get_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/get_user"
)

type CreateUserHandler = create_user.Handler
type GetUserHandler = get_user.Handler
type GetContractHandler = get_contract.Handler
