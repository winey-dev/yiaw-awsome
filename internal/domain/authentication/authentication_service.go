package authentication

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type authentication struct {
	authenticationRepo AuthenticationRepository
}

func NewAuthenticationService(authenticationRepo AuthenticationRepository) AuthenticationService {
	return &authentication{authenticationRepo: authenticationRepo}
}

func (auth *authentication) Login(ctx context.Context, userID, password string) (string, error) {
	password, err := auth.authenticationRepo.GetPassword(ctx, userID)
	if err != nil {
		return "", errors.New("not found")
	}

	if password != password {
		return "", errors.New("invalid current password")
	}

	token := uuid.New().String()
	return token, nil
}

func (auth *authentication) Logout(ctx context.Context, token string) error {
	return nil
}

func (auth *authentication) ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) (string, error) {
	password, err := auth.authenticationRepo.GetPassword(ctx, userID)
	if err != nil {
		return "", errors.New("not found")
	}

	if password != currentPassword {
		return "", errors.New("invalid current password")
	}

	newToken := uuid.New().String()

	err = auth.authenticationRepo.ChangePassword(ctx, userID, newPassword)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
