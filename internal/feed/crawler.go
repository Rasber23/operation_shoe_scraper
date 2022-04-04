package feed

import "C"
import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type brandPaths []string

type productPaths [][]string

func crawl() {

}

func ProductPaths(bp brandPaths) productPaths {

	c := colly.NewCollector()

	var listOfLinks productPaths

	fmt.Println("Längd: ", len(bp))

	c.OnHTML("a[class='JT3_zV CKDt_l ONArL- _2dqvZS lfPP-F CKDt_l LyRfpJ']", func(e *colly.HTMLElement) {

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		link := strings.Fields(e.Attr("href"))

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

	return nil
}

func Paths() brandPaths {

	domain := "https://www.zalando.se/maerken/herrskor/"

	c := colly.NewCollector()

	var listOfLinks brandPaths
	c.OnHTML("a[class='xL5R3s JT3_zV CKDt_l LyRfpJ pVrzNP aX2-iv PO9Fsd _0Hsc7j']", func(e *colly.HTMLElement) {

		listOfLinks = append(listOfLinks, strings.Join(strings.Fields(e.Attr("href")), " "))

	})

	c.Visit(domain)
	for plats, i := range listOfLinks {
		fmt.Println("Link to: ", i+" plats: ", plats)
	}

	fmt.Println(len(listOfLinks))

	return listOfLinks
}
