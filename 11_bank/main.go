package main

import (
	"fmt"
	"github/go_bank_dictionary/bank/accounts"
	"log"
)

func main() {
	account := accounts.CreateAccount("regent0ro")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account)

	fmt.Println(account.Balance())

	err := account.Withdraw(5)
	if err != nil { //if error
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	account.ChangeOwner("hogehoge")
	fmt.Println(account.Onwer())
	fmt.Println(account)
}
