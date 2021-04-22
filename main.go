package main

import (
	accounts "NomadBank/banking"
	"fmt"
)

func main() {
	account := accounts.NewAccount("Keravi")
	account.Desposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account.String())
}
