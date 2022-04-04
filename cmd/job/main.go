package main

import (
	"shoe_scraper/internal/feed"
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

	feed.ProductPaths(feed.Paths())

}
