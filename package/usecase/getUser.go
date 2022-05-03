package usecase

import "github.com/skinnykaen/robbo_control_acces.git/package/gateway"

type GetUserUseCaseImpl struct {
	gateway.Gateway
}

func (s *GetUserUseCaseImpl) Invoke(id uint) error {
	return s.Gateway.GetUser(id)
}
