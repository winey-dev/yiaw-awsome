package account

import (
	"context"
	"time"
)

type Account struct {
	UserID             string
	Password           string
	Email              string
	CreateTime         time.Time
	PasswordUpdateTime time.Time
}

type AccountService interface {
	List(ctx context.Context) ([]*Account, error)
	Get(ctx context.Context, userID string) (*Account, error)
	Create(ctx context.Context, account *Account) (*Account, error)
	Update(ctx context.Context, account *Account) (*Account, error)
	Delete(ctx context.Context, userID string) error
}

type AccountRepository interface {
	List(ctx context.Context) ([]*Account, error)
	Get(ctx context.Context, userID string) (*Account, error)
	Create(ctx context.Context, account *Account) (*Account, error)
	Update(ctx context.Context, account *Account) (*Account, error)
	Delete(ctx context.Context, userID string) error
}
