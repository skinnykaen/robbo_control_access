package gateway

type AuthGatewayImpl struct {
	postgresClient *PostgresClient
}

func (r *AuthGatewayImpl) GetUser(id uint) (err error) {
	// TODO implement me
	return nil
}
