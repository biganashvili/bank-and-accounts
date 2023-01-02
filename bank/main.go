package main

import (
	"log"

	"github.com/biganashvili/bank-and-accounts/bank/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/account", handler.GetAccounts)
	route.POST("/account", handler.CreateAccount)
	route.POST("/account/generate-address", handler.GenerateAddress)
	route.POST("/account/deposit", handler.Deposit)
	route.POST("/account/withdrawal", handler.Withdrawal)
	route.Run(":8085")
	log.Printf("finished")
}
