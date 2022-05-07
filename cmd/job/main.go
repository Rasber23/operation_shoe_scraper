package main

import "github.com/Rasber23/operation_shoe_scraper/internal/chromedp"

type article struct {
	category      string
	brand         string
	articleNumber string
	articleName   string
	price         string
	outerMaterial string
	lining        string
	sole          string
}

type crawlerConfig struct {
	startURL string
}

type brandConfig struct {
	linkElement string
}

func main() {

	/*brand := feed.ScrapeBrand{}
	products := feed.ScrapeProducts{}

	job := job.Crawlers{
		products,
		brand,
	}

	job.Run()*/
	chromedp.ChromedpScraper()
	/*	feed.Scrape()*/
	/*	geziyor.RunGeyz()
	 */
}
