package main

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("element-selector", func(e *colly.HTMLElement) {
		log.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL, r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error on %s: %s", r.Request.URL, err)
	})

	c.Visit("https://scrapeme.live/shop/")
}
