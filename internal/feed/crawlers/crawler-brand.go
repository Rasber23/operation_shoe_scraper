package crawlers

import "C"
import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type ScrapeBrand struct{}

type BrandPaths []string

func (f ScrapeBrand) GetBrandPaths() BrandPaths {

	domain := "https://www.zalando.se/maerken/herrskor/"

	c := colly.NewCollector()

	var listOfLinks BrandPaths

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
