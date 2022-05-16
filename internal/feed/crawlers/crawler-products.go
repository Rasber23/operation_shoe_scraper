package crawlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type ScrapeProducts struct{}

type ProductPaths [][]string

func (f ScrapeProducts) GetProductPaths(bp BrandPaths) ProductPaths {
	fmt.Printf("starting Products: ")

	c := colly.NewCollector()

	var listOfLinks ProductPaths

	fmt.Println("Längd: ", len(bp))

	c.OnHTML("a[class='JT3_zV CKDt_l ONArL- _2dqvZS CKDt_l LyRfpJ']", func(e *colly.HTMLElement) {

		link := strings.Fields(e.Attr("href"))
		fmt.Println(link)
		for _, productLink := range link {
			fmt.Println("test: " + productLink)
		}

		listOfLinks = append(listOfLinks, link)

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL)
	})

	for _, j := range bp {
		c.Visit(j)
	}

	fmt.Print(listOfLinks)

	return listOfLinks
}
