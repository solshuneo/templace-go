package auth

import (
	"errors"
	"lotesaleagent/model"
	"lotesaleagent/model/token"
	"maps"
	"sync"
)

type UserInterface interface {
	Create(user *model.User) model.WrapError
	Find(user *model.User) (*model.User, model.WrapError)
	FindById(userId string) (*model.User, model.WrapError)
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
		Access:  token.Create(maps.Clone(payload), token.ExpiryDurationAccess),
		Refresh: token.Create(maps.Clone(payload), token.ExpiryDurationRefresh),
	}, nil
}

func (auth *Service) GetProfileById(userId string) (*model.User, model.WrapError) {
	var foundUser, err = auth.UserInterface.FindById(userId)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (auth *Service) RefreshToken(refreshToken string) (*token.Token, model.WrapError) {
	var valid, payload = token.VerifyAndDecode(refreshToken)
	if valid == false {
		return nil, model.NewError(errors.New("token is invalid"))
	}
	return &token.Token{
		Access: token.Create(payload, token.ExpiryDurationAccess),
	}, nil
}
