package handler

import (
	"context"
	"errors"
	"io"
	"log"

	pb "github.com/biganashvili/bank-and-accounts/proto_files"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	// dial server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client = pb.NewAccountServiceClient(conn)
}

var client pb.AccountServiceClient

func CreateAccount(c *gin.Context) {

	acc, err := client.CreateAccount(context.Background(), &pb.CreateAccountParams{})
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
	}
	c.JSON(200, acc)
}

type generateAddressReq struct {
	ID string `form:"id"`
}

func GenerateAddress(c *gin.Context) {
	req := &generateAddressReq{}
	if c.ShouldBind(&req) != nil {
		c.JSON(500, errors.New("unxpected amount"))
	}
	acc, err := client.GenerateAddress(context.Background(), &pb.GenerateAddressParams{Id: req.ID})
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
	}
	c.JSON(200, acc)
}

type depositReq struct {
	ID     string `form:"id"`
	Amount uint64 `form:"amount"`
}

func Deposit(c *gin.Context) {
	req := &depositReq{}
	if c.ShouldBind(&req) != nil {
		c.JSON(500, errors.New("unxpected amount"))
	}
	acc, err := client.Deposit(context.Background(), &pb.DepositParams{Id: req.ID, Amount: req.Amount})
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
	}
	c.JSON(200, acc)
}

type withdrawalReq struct {
	ID     string `form:"id"`
	Amount uint64 `form:"amount"`
}

func Withdrawal(c *gin.Context) {
	req := &withdrawalReq{}
	if c.ShouldBind(&req) != nil {
		c.JSON(500, errors.New("unxpected amount"))
	}
	acc, err := client.Withdrawal(context.Background(), &pb.WithdrawalParams{Id: req.ID, Amount: req.Amount})
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
	}
	c.JSON(200, acc)
}

func GetAccounts(c *gin.Context) {
	in := &pb.GetAccountsParams{}
	stream, err := client.GetAccounts(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	for {
		ok := c.Stream(func(w io.Writer) bool {
			resp, err := stream.Recv()
			if err == io.EOF {
				return false
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
				return false
			}
			c.SSEvent("message", resp)
			return true
		})
		if !ok {
			break
		}
	}
}
