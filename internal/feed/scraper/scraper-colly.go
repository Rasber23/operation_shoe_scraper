package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Scrape() {
	fmt.Println("start")
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[class='DT5BTM VHXqc_ rceRmQ _4NtqZU mIlIve']", func(element *colly.HTMLElement) {

		/*price := element.ChildText("p[class='RYghuO uqkIZw ka2E9k uMhVZi FxZV-M pVrzNP']")
		result2 := element.ChildText("div[class='OXMl2N xVaPPo f2qidi _56Chwa']")
		//status := element.ChildText("label[class='_6G4BGa RYghuO _7Cm1F9 ka2E9k uMhVZi FxZV-M sgIQH4 OXMl2N oW5DLR JT3_zV ZkIJC- Md_Vex JCuRr_ UTtITa _0xLoFW FCIprz pVrzNP KRmOLG NuVH8Q']")
		size := element.ChildText("span[class='RYghuO _7Cm1F9 ka2E9k uMhVZi dgII7d pVrzNP']")
		reminder := element.ChildText("span[class='RYghuO u-6V88 ka2E9k uMhVZi dgII7d wu3klO']")
		fewLeft := element.ChildText("span[class='RYghuO u-6V88 ka2E9k uMhVZi FxZV-M pVrzNP']")
		fmt.Println("Result 1: ", price)

		fmt.Println("result 2", result2)

		fmt.Println("-------------------------------------")

		fmt.Println("size 3: ", size)
		fmt.Println("påMINMAJ 3: ", reminder)
		fmt.Println("få kvar: ", fewLeft)*/

		/*re := regexp.MustCompile(`[A-Z][^A-Z]*`)
		submatchall := re.FindAllString(result2, -1)
		for i, s := range submatchall {
			fmt.Println("index: ", i, s)
		}*/
		contentet := element.Text

		fmt.Println(contentet)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.Visit("https://www.zalando.se/nike-sportswear-air-force-1-sneakers-whiteblack-ni114d0i0-a11.html")
}
