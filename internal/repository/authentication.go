package repository

import (
	"context"
	"errors"
	"yiaw-awsome/internal/domain/authentication"
)

type authenticationRepo struct{}

func NewAuthenticationRepository() authentication.AuthenticationRepository {
	return &authenticationRepo{}
}

func (authRepo *authenticationRepo) GetPassword(ctx context.Context, userID string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	v, ok := etcd[userID]
	if !ok {
		return "", errors.New("not found")
	}
	return v.Password, nil
}

func (authRepo *authenticationRepo) ChangePassword(ctx context.Context, userID, newPassword string) error {
	mutex.Lock()
	defer mutex.Unlock()

	v, ok := etcd[userID]
	if !ok {
		return errors.New("not found")
	}

	v.Password = newPassword
	return nil
}
