package api

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user"
)

func InitModule() module.Set {
	return module.Set{
		User:     user.New(),
		Contract: contract.New(),
	}
}
