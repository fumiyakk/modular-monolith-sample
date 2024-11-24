package registry

import "github.com/fumiyakk/modular-monolith-sample/internal/server/module"

type Registry interface {
	Module() module.Set
}
