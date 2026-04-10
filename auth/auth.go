package auth

import (
	"lotesaleagent/model"
	"sync"
)

type UserInterface interface {
	Create(user *model.User) model.WrapError
	Find(user *model.User) model.WrapError
}

type AuthService struct {
	UserInterface UserInterface
}

func (auth *AuthService) Register(user *model.User) (err model.WrapError) {
	mutex := sync.Mutex{}
	mutex.Lock()
	err = auth.UserInterface.Create(user)
	mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func (auth *AuthService) Login(user *model.User) (token *model.Token, err model.WrapError) {
	err = auth.UserInterface.Find(user)
	return model.NewToken(), err
}
