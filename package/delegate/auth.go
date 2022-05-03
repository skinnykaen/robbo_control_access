package delegate

import "github.com/skinnykaen/robbo_control_acces.git/package/usecase"

type AuthDelegate struct {
	usecase.GetUserUseCase
}

func (s *AuthDelegate) GetUser(id uint) (err error) {
	return nil
}
