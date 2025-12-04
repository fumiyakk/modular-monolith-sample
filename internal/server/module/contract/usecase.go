package contract

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/usecase/create_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/usecase/get_contract"
)

type CreateContractUsecase = create_contract.Usecase
type GetContractUsecase = get_contract.Usecase
