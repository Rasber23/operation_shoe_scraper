package main

import (
	"shoe_scraper/internal/feed"
	"shoe_scraper/internal/job"
)

type article struct {
	categorie     string
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

	brand := feed.ScrapeBrand{}
	products := feed.ScrapeProducts{}

	job := job.Crawlers{
		products,
		brand,
	}

	job.Run()

}
