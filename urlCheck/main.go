package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

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

	//results := map[string]string{} //맵은 이렇게 만들거나
	results := make(map[string]string) //이런식으로 초기화를 해줘야 값을 넣을 수 있음
	c := make(chan requestResult)
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		//fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	// c := make(chan string)
	// people := [4]string{"chobkyu", "bk", "cho", "minsu"}
	// for _, person := range people {
	// 	go isSexy(person, c)
	// }

	// for i := 0; i < len(people); i++ {
	// 	fmt.Println(<-c)
	// }

	//time.Sleep(time.Second * 5)
}

func hitURL(url string, c chan<- requestResult) { //<- 표시로 채널에서 데이터를 받을 수는 없고 보낼 수만 있다고 지정하는 것
	fmt.Println("checking : ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

func isSexy(person string, c chan string) {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(person, "is sexy", i)
	// 	time.Sleep(time.Second)
	// }
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
