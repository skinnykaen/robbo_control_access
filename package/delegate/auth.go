package delegate

import "github.com/skinnykaen/robbo_control_acces.git/package/usecase"

type AuthDelegate struct {
	usecase.AuthUseCase
}

func (s *AuthDelegate) SignIn(email, password string) (err error) {
	return nil
}
