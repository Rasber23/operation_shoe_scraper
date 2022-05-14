package main

import (
	"github.com/Rasber23/operation_shoe_scraper/internal/bigquery"
	"github.com/Rasber23/operation_shoe_scraper/internal/business/job"
	dom_parser "github.com/Rasber23/operation_shoe_scraper/internal/dom-parser"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/crawlers"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/scraper"
)

type crawlerConfig struct {
	startURL string
}

type brandConfig struct {
	linkElement string
}

func main() {

	brand := crawlers.ScrapeBrand{}
	products := crawlers.ScrapeProducts{}

	scraper := scraper.Scraper{}
	transformer := dom_parser.Transform{}
	saver := bigquery.BQService{
		ProjectId: "shoe-crawler",
		DataSetId: "zalando",
		TableId:   "articles",
	}

	job := job.Crawlers{
		Scrape:        scraper,
		Transform:     transformer,
		CrawlProducts: products,
		CrawlBrands:   brand,
		Saver:         saver,
	}

	job.Run()

	//scraper.Run()
	/*	feed.Scrape()*/
	/*	geziyor.RunGeyz()
	 */
}
