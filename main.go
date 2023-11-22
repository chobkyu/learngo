package main

import (
	"fmt"
	"learngo/github.com/serranoarevalo/learngo/accounts"
)

func main() {
	account := accounts.NewAccout("chobkyu")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err) //에러나면 프로그램 종료
		fmt.Println(err)
	}
	fmt.Println(account.String())
}
