package delegate

import (
	"github.com/skinnykaen/robbo_control_acces.git/package/usecase"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out
	AuthDelegate
}

func Setup(getUserUseCase usecase.GetUserUseCase) Module {
	return Module{
		AuthDelegate: AuthDelegate{
			getUserUseCase,
		},
	}
}
