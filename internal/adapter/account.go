package adapter

import (
	"context"
	"yiaw-awsome/internal/domain/account"
	v1accountpb "yiaw-awsome/proto/v1/account"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Account struct {
	v1accountpb.AccountServiceServer
	accountService account.AccountService
}

func NewAccount(accountService account.AccountService) *Account {
	return &Account{accountService: accountService}
}

func (a *Account) ListAccounts(ctx context.Context, req *v1accountpb.ListAccountsRequest) (*v1accountpb.ListAccountsResponse, error) {
	accounts, err := a.accountService.List(ctx)
	if err != nil {
		return &v1accountpb.ListAccountsResponse{}, status.Errorf(codes.NotFound, "accounts is empty")
	}

	response := &v1accountpb.ListAccountsResponse{}
	for _, account := range accounts {
		response.Accounts = append(
			response.Accounts,
			&v1accountpb.Account{
				UserId:             account.UserID,
				Email:              account.Email,
				CreateTime:         timestamppb.New(account.CreateTime),
				PasswordUpdateTime: timestamppb.New(account.PasswordUpdateTime),
			},
		)
	}
	return response, nil
}

func (a *Account) GetAccount(ctx context.Context, req *v1accountpb.GetAccountRequest) (*v1accountpb.Account, error) {

	account, err := a.accountService.Get(ctx, req.UserId)
	if err != nil {
		return &v1accountpb.Account{}, status.Errorf(codes.NotFound, "not found %s", req.UserId)
	}

	return &v1accountpb.Account{
		UserId:             account.UserID,
		Email:              account.Email,
		CreateTime:         timestamppb.New(account.CreateTime),
		PasswordUpdateTime: timestamppb.New(account.PasswordUpdateTime),
	}, nil
}

func (a *Account) CreateAccount(ctx context.Context, req *v1accountpb.CreateAccountRequest) (*v1accountpb.Account, error) {
	account, err := a.accountService.Create(
		ctx,
		&account.Account{
			UserID:   req.UserId,
			Password: req.Password,
			Email:    req.Email,
		},
	)

	if err != nil {
		return &v1accountpb.Account{}, status.Errorf(codes.AlreadyExists, "%s already exists", req.UserId)
	}
	return &v1accountpb.Account{
		UserId:             account.UserID,
		Email:              account.Email,
		CreateTime:         timestamppb.New(account.CreateTime),
		PasswordUpdateTime: timestamppb.New(account.PasswordUpdateTime),
	}, nil
}

func (a *Account) UpdateAccount(ctx context.Context, req *v1accountpb.UpdateAccountRequest) (*v1accountpb.Account, error) {
	account, err := a.accountService.Update(
		ctx,
		&account.Account{
			UserID: req.UserId,
			Email:  req.Email,
		},
	)

	if err != nil {
		return &v1accountpb.Account{}, status.Errorf(codes.NotFound, "%s not found", req.UserId)
	}
	return &v1accountpb.Account{
		UserId:             account.UserID,
		Email:              account.Email,
		CreateTime:         timestamppb.New(account.CreateTime),
		PasswordUpdateTime: timestamppb.New(account.PasswordUpdateTime),
	}, nil
}

func (a *Account) DeleteAccount(ctx context.Context, req *v1accountpb.DeleteAccountRequest) (*emptypb.Empty, error) {
	err := a.accountService.Delete(ctx, req.UserId)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "unexpect internal server error")
	}
	return &emptypb.Empty{}, nil
}
