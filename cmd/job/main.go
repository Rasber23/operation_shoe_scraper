package main

import (
	"fmt"
	"github.com/Rasber23/operation_shoe_scraper/internal/bigquery"
	"github.com/Rasber23/operation_shoe_scraper/internal/business/job"
	dom_parser "github.com/Rasber23/operation_shoe_scraper/internal/dom-parser"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/crawlers"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/scraper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

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

	port := os.Getenv("PORT")
	port = fmt.Sprintf(":%v", port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		job.Run()
	})

	er := http.ListenAndServe(port, r)
	if er != nil {
		fmt.Errorf("%v", er)
	}
}
