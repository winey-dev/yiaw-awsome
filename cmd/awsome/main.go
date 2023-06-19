package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	config "yiaw-awsome/config/awsome"
	"yiaw-awsome/internal/adapter"
	"yiaw-awsome/internal/domain/account"
	"yiaw-awsome/internal/domain/authentication"
	"yiaw-awsome/internal/repository"
	v1accountpb "yiaw-awsome/proto/v1/account"
	v1authenticationpb "yiaw-awsome/proto/v1/authentication"
)

func main() {

	cfg := config.LoadConfiguration()

	accountService := account.NewAccountService(repository.NewAccountRepository())
	authenticationService := authentication.NewAuthenticationService(repository.NewAuthenticationRepository())

	listen, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%s", cfg.Address, cfg.GRPCPort),
	)

	if err != nil {
		fmt.Printf("socket listen failed. endpoint=%s:%s, err=%v\n", cfg.Address, cfg.GRPCPort, err)
		return
	}

	server := grpc.NewServer()
	v1accountpb.RegisterAccountServiceServer(server, adapter.NewAccount(accountService))
	v1authenticationpb.RegisterAuthenticationServiceServer(server, adapter.NewAuthentication(authenticationService))
	reflection.Register(server)

	fmt.Printf("Start Server %s:%s\n", cfg.Address, cfg.GRPCPort)
	server.Serve(listen)
	defer server.Stop()
	return
}
