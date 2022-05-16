package scraper

import (
	"context"
	"fmt"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/crawlers"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
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
		//chromedp.WithDebugf(log.Printf),
	)

	defer cancel()

	var doms ArticleDoms

	for i, urlBrand := range productUrls {

		for _, productUrls := range urlBrand {
			fmt.Printf("Visiting %v", productUrls)

			r := fetch(ctx, productUrls)
			doms = append(doms, r)
		}
		if i == 5 {
			break
		}
	}
	return doms
}

func fetch(ctx context.Context, productUrls string) ArticleDom {

	fmt.Println("I context link: " + productUrls)
	var value string
	var dom ArticleDom

	err := chromedp.Run(ctx,

		chromedp.Emulate(device.IPad),

		chromedp.Navigate(productUrls),

		chromedp.OuterHTML(`/html/body/div[4]/div/div[1]/div/div/div[2]`, &dom.ProductInf, chromedp.NodeVisible),

		chromedp.Text(`/html/body/div[4]/div/div[1]/div/div/div[2]/div[1]/x-wrapper-re-1-7/div/div[1]/div/div/button/span/span`, &value),
	)
	if err != nil {
		log.Fatal(err)
	}

	if value == "Välj storlek" {
		err := chromedp.Run(ctx,
			chromedp.Navigate(productUrls),
			chromedp.Click(`#picker-trigger`, chromedp.NodeVisible),
			chromedp.OuterHTML(`div[class='F8If-J mZoZK2 KLaowZ hj1pfK _8n7CyI JCuRr_']`, &dom.SizeInf, chromedp.NodeVisible),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	dom.SizeInf = value

	return dom
}
