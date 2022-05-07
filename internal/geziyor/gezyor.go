package geziyor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

/*func RunGeyz() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
            g.Do()
			g.GetRendered("https://www.zalando.se/ecco-x-trinsic-m-vandringssandaler-black-ec112g00v-q11.html", g.Opt.ParseFunc)

		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			test := r.

			element := r.HTMLDoc.Find("span[class='RYghuO _7Cm1F9 ka2E9k uMhVZi dgII7d D--idb']")

			fmt.Println(element)
		},
	}).Start()
}*/

func quotesParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"text":   s.Find("span.text").Text(),
			"author": s.Find("small.author").Text(),
		}
	})
	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
		g.Get(r.JoinURL(href), quotesParse)
	}
}
