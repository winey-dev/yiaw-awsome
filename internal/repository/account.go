package repository

import (
	"context"
	"errors"
	"sync"
	"time"
	"yiaw-awsome/internal/domain/account"
)

var mutex sync.Mutex
var etcd map[string]*account.Account

func init() {
	etcd = make(map[string]*account.Account)
	etcd["admin"] = &account.Account{
		UserID:             "admin",
		Password:           "admin1234",
		Email:              "admin@example.com",
		CreateTime:         time.Now(),
		PasswordUpdateTime: time.Now(),
	}
	etcd["lee"] = &account.Account{
		UserID:             "lee",
		Password:           "lee1234",
		Email:              "lee@example.com",
		CreateTime:         time.Now(),
		PasswordUpdateTime: time.Now(),
	}
}

type accountRepo struct{}

func NewAccountRepository() account.AccountRepository {
	return &accountRepo{}
}

func (ar *accountRepo) List(ctx context.Context) ([]*account.Account, error) {
	var accounts []*account.Account

	mutex.Lock()
	defer mutex.Unlock()

	for _, v := range etcd {
		accounts = append(
			accounts,
			&account.Account{
				UserID:             v.UserID,
				Password:           v.Password,
				Email:              v.Email,
				CreateTime:         v.CreateTime,
				PasswordUpdateTime: v.PasswordUpdateTime,
			},
		)
	}

	if len(accounts) == 0 {
		return nil, errors.New("not found")
	}
	return accounts, nil
}

func (ar *accountRepo) Get(ctx context.Context, userID string) (*account.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	v, ok := etcd[userID]
	if !ok {
		return nil, errors.New("not found")
	}
	return &account.Account{
		UserID:             v.UserID,
		Password:           v.Password,
		Email:              v.Email,
		CreateTime:         v.CreateTime,
		PasswordUpdateTime: v.PasswordUpdateTime,
	}, nil
}

func (ar *accountRepo) Create(ctx context.Context, createAccount *account.Account) (*account.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := etcd[createAccount.UserID]
	if ok {
		return nil, errors.New("already exists")
	}

	v := &account.Account{
		UserID:             createAccount.UserID,
		Password:           createAccount.Password,
		Email:              createAccount.Email,
		CreateTime:         time.Now(),
		PasswordUpdateTime: time.Now(),
	}
	etcd[createAccount.UserID] = v
	return v, nil
}

func (ar *accountRepo) Update(ctx context.Context, updateAccount *account.Account) (*account.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()
	v, ok := etcd[updateAccount.UserID]
	if !ok {
		return nil, errors.New("not found")
	}

	v.Email = updateAccount.Email

	return v, nil
}

func (ar *accountRepo) Delete(ctx context.Context, userID string) error {
	mutex.Lock()
	defer mutex.Unlock()
	delete(etcd, userID)
	return nil
}
