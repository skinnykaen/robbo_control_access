package gateway

import "go.uber.org/fx"

type Gateway interface {
	GetUser(id uint) (err error)
}

type Module struct {
	fx.Out
	Gateway
}

func Setup(postgresClient PostgresClient) Module {
	return Module{
		Gateway: &AuthGatewayImpl{postgresClient: &postgresClient},
	}
}
