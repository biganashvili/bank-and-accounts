package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/biganashvili/bank-and-accounts/db"
	pb "github.com/biganashvili/bank-and-accounts/proto_files"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server
type server struct {
	pb.UnimplementedAccountServiceServer
}

// Deposit implements
func (s *server) Deposit(ctx context.Context, in *pb.DepositParams) (*pb.Account, error) {
	account, err := db.GetAccount(in.Id)

	if err != nil {
		return nil, err
	}
	account.Balance += in.Amount
	err = db.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// Withdrawal implements
func (s *server) Withdrawal(ctx context.Context, in *pb.WithdrawalParams) (*pb.Account, error) {
	account, err := db.GetAccount(in.Id)

	if err != nil {
		return nil, err
	}

	if account.Balance < in.Amount {
		return nil, errors.New("insufficient funds")
	}
	account.Balance -= in.Amount
	err = db.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// GenerateAddress implements
func (s *server) GenerateAddress(ctx context.Context, in *pb.GenerateAddressParams) (*pb.Account, error) {
	account, err := db.GetAccount(in.Id)

	if err != nil {
		return nil, err
	}
	newUUID := uuid.New()
	account.WalletID = newUUID.String()
	err = db.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// CreateAccount implements
func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountParams) (*pb.Account, error) {

	newUUID := uuid.New()
	account := pb.Account{Id: newUUID.String(), Balance: 0, WalletID: ""}
	id, err := db.InsertAccount(&account)
	log.Println(id)
	if err != nil {
		return &pb.Account{}, err
	}
	return &account, nil
}

// GetAccounts implements
func (s *server) GetAccounts(in *pb.GetAccountsParams, srv pb.AccountService_GetAccountsServer) error {
	accounts, err := db.GetAllAccounts()
	if err != nil {
		return err
	}
	for _, account := range accounts {
		err := srv.Send(account)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
