package job

import (
	"fmt"
	"github.com/Rasber23/operation_shoe_scraper/internal/business/entity"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/crawlers"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/scraper"
)

type Crawlers struct {
	Scrape        IScraper
	Transform     ITransform
	CrawlProducts ICrawlProducts
	CrawlBrands   ICrawlBrands
	Saver         IPersist
}

type IPersist interface {
	SaveArticles(articles entity.Articles)
}

type IScraper interface {
	Run(productUrls crawlers.ProductPaths) scraper.ArticleDoms
}
type ITransform interface {
	Run(domes scraper.ArticleDoms) entity.Articles
}

type ICrawlProducts interface {
	GetProductPaths(bp crawlers.BrandPaths) crawlers.ProductPaths
}

type ICrawlBrands interface {
	GetBrandPaths() crawlers.BrandPaths
}

func (crawler Crawlers) Run() {

	brandUrls := crawler.CrawlBrands.GetBrandPaths()

	productUrls := crawler.CrawlProducts.GetProductPaths(brandUrls)

	domsHtml := crawler.Scrape.Run(productUrls)

	arts := crawler.Transform.Run(domsHtml)

	//crawler.Saver.SaveArticles(arts)

	for _, art := range arts {
		fmt.Println(art)
	}

}
