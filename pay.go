package main

import (
	"cielo/transactions"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	jsonstring := transactions.AuthorizeCreditCard(
		"1234567890",
		250,
		"4551870000000181",
		"Teste Holder",
		"12/2021",
		"123",
		"Visa",
	)
	fmt.Println(jsonstring)
}
