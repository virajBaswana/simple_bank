package db

import (
	"context"
	"fmt"
	db "simple_bank/db/sqlc"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	args := db.CreateAccountParams{
		Owner:    "Viraj",
		Balance:  10000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	if err != nil {
		fmt.Println("Error while creating a new account : %s", err)
	}
}
