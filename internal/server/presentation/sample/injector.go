package sample

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/get_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/presentation/sample/internal/get_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario"
)

type HandlerInjector interface {
	CreateUserHandler
	GetUserHandler
	GetContractHandler
}

func New(s scenario.ScenarioInjector) HandlerInjector {
	return &struct {
		CreateUserHandler
		GetUserHandler
		GetContractHandler
	}{
		create_user.New(s),
		get_user.New(s),
		get_contract.New(s),
	}
}
