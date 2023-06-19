package authentication

import "context"

type AuthenticationService interface {
	Login(ctx context.Context, userID, password string) (string, error)
	Logout(ctx context.Context, token string) error
	ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) (string, error)
}

type AuthenticationRepository interface {
	GetPassword(ctx context.Context, userID string) (string, error)
	ChangePassword(ctx context.Context, userID, newPassword string) error
}
