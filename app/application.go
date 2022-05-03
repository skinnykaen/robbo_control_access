package app

import (
	"github.com/skinnykaen/robbo_control_acces.git/package/config"
	"github.com/skinnykaen/robbo_control_acces.git/package/gateway"
	"github.com/skinnykaen/robbo_control_acces.git/package/logger"
	"go.uber.org/fx"
)

func AppInvokeWith(options ...fx.Option) *fx.App {
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(gateway.Setup),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	AppInvokeWith(fx.Invoke()).Run()
}
