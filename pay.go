package main

import (
	"cielo/transactions"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	jsonstring := transactions.AuthorizeCreditCard()
	fmt.Println(jsonstring)
}
