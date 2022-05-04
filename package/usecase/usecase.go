package usecase

import (
	"github.com/skinnykaen/robbo_control_acces.git/package/gateway"
	"go.uber.org/fx"
)

type AuthUseCase interface {
	GetUser(email, password string) (err error)
}

type Module struct {
	fx.Out
	AuthUseCase
}

func Setup(gateway gateway.AuthGateway) Module {
	return Module{
		AuthUseCase: &AuthUseCaseImpl{gateway},
	}
}
