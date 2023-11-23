package main

import (
	"fmt"
	"learngo/github.com/serranoarevalo/learngo/myDict"
)

//account ex
// func main() {
// 	account := accounts.NewAccout("chobkyu")
// 	account.Deposit(10)
// 	fmt.Println(account.Balance())
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		//log.Fatalln(err) //에러나면 프로그램 종료
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.String())
// }

func main() {
	// dictionary := myDict.Dictionary{"first": "firstword"}
	// definition, err := dictionary.Search("first")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	dictionary := myDict.Dictionary{}
	err := dictionary.Add("hi", "hello")

	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search("hi")

	fmt.Println(definition)

	err2 := dictionary.Update("dsfsadf", "bye")
	if err2 != nil {
		fmt.Println(err2)
	}

	word, _ := dictionary.Search("hi")
	fmt.Println(word)

	testWord, _ := dictionary.Search("hi")
	dictionary.Delete("hi")

	check, err := dictionary.Search("hi")
	fmt.Println(testWord)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(check)

}
