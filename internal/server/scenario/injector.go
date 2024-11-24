package scenario

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/lib/unit_of_work"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/get_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/get_user"
)

type ScenarioInjector interface {
	CreateUserScenario
	GetUserScenario
	GetContractScenario
}

func New(uow unit_of_work.UnitOfWork, m module.Set) ScenarioInjector {
	return &struct {
		CreateUserScenario
		GetUserScenario
		GetContractScenario
	}{
		create_user.New(uow, m.User, m.Contract),
		get_user.New(uow, m.User),
		get_contract.New(uow, m.Contract),
	}
}
