package chromedp

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"strings"
	"time"
)

func ChromedpScraper() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	/*	ctx, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...)
		defer cancel()
		ctx, cancel = chromedp.NewContext(ctx,
			chromedp.WithDebugf(log.Printf))
		defer cancel()*/

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var res string
	var res2 string
	urls := []string{
		"https://www.zalando.se/ecco-x-trinsic-m-vandringssandaler-black-ec112g00v-q11.html",
		"https://www.zalando.se/nike-sportswear-air-force-1-sneakers-whiteblack-ni112o0h3-a11.html",
	}

	for _, url := range urls {

	}

	err := chromedp.Run(ctx,

		chromedp.Emulate(device.IPad),
		chromedp.Navigate(`https://www.zalando.se/ecco-x-trinsic-m-vandringssandaler-black-ec112g00v-q11.html`),

		// find and click "Example" link
		//chromedp.Click(`button[class='J1Rmt- _6-WsK3 Md_Vex Nk_Omi _MmCDa _0xLoFW FCIprz NN8L-8 _7Cm1F9 ka2E9k uMhVZi FxZV-M ovwSlD LyRfpJ K82if3 heWLCX mo6ZnF pVrzNP']`, chromedp.NodeVisible),
		chromedp.Click(`#picker-trigger`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		/*		chromedp.OuterHTML(`span[class='RYghuO _7Cm1F9 ka2E9k uMhVZi dgII7d D--idb']`, &res, chromedp.NodeVisible),
		 */
		chromedp.OuterHTML(`div[class='F8If-J mZoZK2 KLaowZ hj1pfK _8n7CyI JCuRr_']`, &res, chromedp.NodeVisible),

		//chromedp.Value(`span[class='RYghuO _7Cm1F9 ka2E9k uMhVZi dgII7d pVrzNP']`, &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	p := strings.NewReader(res)
	p2 := strings.NewReader(res2)

	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		log.Fatal(err)
	}
	doc2, err := goquery.NewDocumentFromReader(p2)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div > span").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		size := s.Text()
		fmt.Printf("size %d: %s\n", i, size)
	})

	doc2.Find("div > span").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		size := s.Text()
		fmt.Printf("size2 %d: %s\n", i, size)
	})
}
