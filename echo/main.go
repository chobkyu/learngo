package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := c.FormValue("term")
	fmt.Println(term)

	return nil
}
