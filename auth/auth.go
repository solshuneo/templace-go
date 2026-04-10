package auth

import (
	"lotesaleagent/model"
	"sync"
)

type UserInterface interface {
	Create(user *model.User) model.WrapError
	Find(user *model.User) model.WrapError
}

type Service struct {
	UserInterface UserInterface
}

func (auth *Service) Register(user *model.User) model.WrapError {
	mutex := sync.Mutex{}
	mutex.Lock()
	var err = auth.UserInterface.Create(user)
	mutex.Unlock()
	return err
}

func (auth *Service) Login(user *model.User) (*model.Token, model.WrapError) {
	var err = auth.UserInterface.Find(user)
	return model.NewToken(), err
}
