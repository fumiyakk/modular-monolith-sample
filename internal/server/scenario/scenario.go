package scenario

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/create_user"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/get_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/scenario/internal/get_user"
)

type CreateUserScenario = create_user.Scenario
type GetUserScenario = get_user.Scenario
type GetContractScenario = get_contract.Scenario
