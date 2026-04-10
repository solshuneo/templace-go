package auth

import (
	"lotesaleagent/model"
)

type UserInterface interface {
	Create(user *model.User) model.WrapError
}

type AuthService struct {
	UserInterface UserInterface
}

func (auth *AuthService) Register(user *model.User) (err model.WrapError) {

	err = auth.UserInterface.Create(user)
	return err
}
