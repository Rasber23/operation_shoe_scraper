package feed

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func scrape() {
	fmt.Println("start")
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("Body", func(element *colly.HTMLElement) {

		result1 := element.ChildText("span[class='uqkIZw ka2E9k uMhVZi FxZV-M RYghuO pVrzNP']")
		result2 := element.ChildText(" div[class='OXMl2N xVaPPo f2qidi _56Chwa']")
		result3 := element.ChildText("div[class='mt1kvu ka2E9k uMhVZi FxZV-M SZKKsK pVrzNP _5Yd-hZ']")
		re := regexp.MustCompile(`[A-Z][^A-Z]*`)

		fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

		submatchall := re.FindAllString(result2, -1)

		fmt.Println("Result one: ", result1)
		fmt.Println("Result 3: ", result3)

		for i, s := range submatchall {
			fmt.Println("index: ", i, s)
		}

	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.Visit("https://www.zalando.se/nike-sportswear-air-force-1-sneakers-whiteblack-ni114d0i0-a11.html")
}
