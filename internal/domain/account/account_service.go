package account

import "context"

type account struct {
	accountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) AccountService {
	return &account{accountRepository: accountRepository}
}

func (a *account) List(ctx context.Context) ([]*Account, error) {
	return a.accountRepository.List(ctx)
}

func (a *account) Get(ctx context.Context, userID string) (*Account, error) {
	return a.accountRepository.Get(ctx, userID)
}

func (a *account) Create(ctx context.Context, account *Account) (*Account, error) {
	return a.accountRepository.Create(ctx, account)
}

func (a *account) Update(ctx context.Context, account *Account) (*Account, error) {
	return a.accountRepository.Update(ctx, account)
}

func (a *account) Delete(ctx context.Context, userID string) error {
	return a.accountRepository.Delete(ctx, userID)
}
