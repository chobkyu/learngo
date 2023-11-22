package main

import (
	"fmt" //formatting을 위한 함수
	"learngo/github.com/serranoarevalo/learngo/something"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) { //naked return
	//return len(name), strings.ToUpper(name)
	defer fmt.Println("i'm groot") //함수가 끝나고 나서 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// 반복문
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { //index와 배열요소 반환
		//fmt.Println(number, index)
		total += number
	}
	return total
}

// if문
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

// switch
func canIDringSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// pointer
func pointerEx() {
	a := 2
	b := &a
	fmt.Println(*b)
}

func arrayEx() {
	names := [5]string{"hi", "hello", "bye"} //array
	names[3] = "good"
	names[4] = "bye"
	fmt.Println(names)

	drinks := []string{"soju", "beer"} //slice
	drinks = append(drinks, "wine")
	fmt.Println(drinks)
}

func mapEx() {
	drinks := map[string]string{"name": "soju", "dosu": "16"} //[key]value 형식임
	fmt.Println(drinks)
	for _, value := range drinks {
		fmt.Println(value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	fmt.Println("hello world")
	something.SayHello()
	fmt.Println(multiply(2, 2))
	totalLen, _ := lenAndUpper("chobkyu")
	fmt.Println(totalLen)
	repeatMe("hi", "hello", "mell")
	fmt.Println(lenAndUpper("chobkyu"))
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
	fmt.Println(canIDrink(14))
	fmt.Println(canIDringSwitch(18))
	pointerEx()
	arrayEx()
	mapEx()
	favFood := []string{"meat", "fish"}
	chobkyu := person{name: "bk", age: 26, favFood: favFood}
	fmt.Println(chobkyu)
}
