package usecase

import (
	"github.com/skinnykaen/robbo_control_acces.git/package/gateway"
	"go.uber.org/fx"
)

type GetUserUseCase interface {
	Invoke(id uint) (err error)
}

type Module struct {
	fx.Out
	GetUserUseCase
}

func Setup(gateway gateway.Gateway) Module {
	return Module{
		GetUserUseCase: &GetUserUseCaseImpl{gateway},
	}
}
