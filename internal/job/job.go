package job

import (
	"shoe_scraper/internal/feed"
)

type Crawlers struct {
	CrawlProducts ICrawlProducts
	CrawlBrands   ICrawlBrands
}

type ICrawlProducts interface {
	GetProductPaths(bp feed.BrandPaths) feed.ProductPaths
}

type ICrawlBrands interface {
	GetBrandPaths() feed.BrandPaths
}

func (crawl Crawlers) Run() {

	crawl.CrawlProducts.GetProductPaths(crawl.CrawlBrands.GetBrandPaths())

}
