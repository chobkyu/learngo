package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&l=%EC%84%9C%EC%9A%B8&vjk=701e887d6f3eb831"

func Scrapper() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	req, rErr := http.NewRequest("GET", baseURL, nil)
	checkErr(rErr)

	// 프록시로 호출
	purl, err := url.Parse(baseURL)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}
	res, err := client.Do(req)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	fmt.Println(doc)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)

	}
}
