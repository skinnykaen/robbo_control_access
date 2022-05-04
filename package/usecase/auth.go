package usecase

import "github.com/skinnykaen/robbo_control_acces.git/package/gateway"

type AuthUseCaseImpl struct {
	gateway.AuthGateway
}

func (s *AuthUseCaseImpl) SignIn(email, password string) error {
	return s.AuthGateway.GetUser(email, password)
}
