package main

import (
	"log"

	"github.com/M-Horizon-GmbH/colly-with-req"
)

func main() {
	c := colly.NewCollector()
	c.ImpersonateChrome()

	c.OnHTML("element-selector", func(e *colly.HTMLElement) {
		log.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error on %s: %s", r.Request.URL, err)
	})

	c.Visit("https://www.howsmyssl.com/")
}
