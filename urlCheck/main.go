package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	//result := map[string]string{} //맵은 이렇게 만들거나
	var results = make(map[string]string) //이런식으로 초기화를 해줘야 값을 넣을 수 있음
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)

	}

	c := make(chan string)
	people := [4]string{"chobkyu", "bk", "cho", "minsu"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	//time.Sleep(time.Second * 5)
}

func hitURL(url string) error {
	fmt.Println("checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}

func isSexy(person string, c chan string) {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(person, "is sexy", i)
	// 	time.Sleep(time.Second)
	// }
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
