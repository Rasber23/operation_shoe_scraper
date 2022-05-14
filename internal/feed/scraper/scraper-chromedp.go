package scraper

import (
	"context"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/crawlers"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"time"
)

type ArticleDom struct {
	ProductInf string
	SizeInf    string
}

type Scraper struct{}

type ArticleDoms []ArticleDom

func (scraper Scraper) Run(productUrls crawlers.ProductPaths) ArticleDoms {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//scraper.WithDebugf(log.Printf),
	)
	/*	ctx, cancel := scraper.NewExecAllocator(context.Background(), append(scraper.DefaultExecAllocatorOptions[:], scraper.Flag("headless", false))...)
		defer cancel()
		ctx, cancel = scraper.NewContext(ctx,
			scraper.WithDebugf(log.Printf))
		defer cancel()*/

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var doms ArticleDoms

	for _, urlBrand := range productUrls {
		for _, productUrls := range urlBrand {

			var dom ArticleDom

			err := chromedp.Run(ctx,

				chromedp.Emulate(device.IPad),

				chromedp.Navigate(productUrls),

				chromedp.OuterHTML(`/html/body/div[4]/div/div[1]/div/div/div[2]`, &dom.ProductInf, chromedp.NodeVisible),

				chromedp.Click(`#picker-trigger`, chromedp.NodeVisible),
				chromedp.OuterHTML(`div[class='F8If-J mZoZK2 KLaowZ hj1pfK _8n7CyI JCuRr_']`, &dom.SizeInf, chromedp.NodeVisible),
			)
			if err != nil {
				log.Fatal(err)
			}

			doms = append(doms, dom)
		}
	}
	return doms
}
