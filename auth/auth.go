package auth

import (
	"lotesaleagent/internal/repository/gsql"
	"lotesaleagent/model"
	"lotesaleagent/model/token"
	"sync"
)

type UserInterface interface {
	Create(user *model.User) model.WrapError
	Find(user *model.User) (*gsql.User, model.WrapError)
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

func (auth *Service) Login(user *model.User) (*token.Token, model.WrapError) {
	var foundUser, err = auth.UserInterface.Find(user)
	if err != nil {
		return nil, err
	}
	payload := make(map[string]any)
	payload["id"] = foundUser.Id
	payload["username"] = foundUser.Username
	payload["password"] = foundUser.Password
	return &token.Token{
		Access:  token.Create(make(map[string]any), token.ExpiryDurationAccess),
		Refresh: token.Create(make(map[string]any), token.ExpiryDurationRefresh),
	}, nil
}
