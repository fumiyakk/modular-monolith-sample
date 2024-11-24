package module

import (
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/contract"
	"github.com/fumiyakk/modular-monolith-sample/internal/server/module/user"
)

type Set struct {
	Contract contract.Module
	User     user.Module
}
