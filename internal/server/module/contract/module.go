package contract

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/repository"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/usecase/create_contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract/internal/usecase/get_contract"
)

type Module interface {
	CreateContractUsecase
	GetContractUsecase
}

func New() Module {
	repo := repository.New()

	return &struct {
		CreateContractUsecase
		GetContractUsecase
	}{
		create_contract.New(repo),
		get_contract.New(repo),
	}
}
