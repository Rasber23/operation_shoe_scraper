package dom_parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Rasber23/operation_shoe_scraper/internal/business/entity"
	"github.com/Rasber23/operation_shoe_scraper/internal/feed/scraper"
	"log"
	"strings"
)

type Transform struct{}

func (Transform Transform) Run(domes scraper.ArticleDoms) entity.Articles {

	var articles entity.Articles
	for _, dom := range domes {
		article := transformToArticle(dom)

		articles = append(articles, article)
	}

	return articles
}

func transformToArticle(dom scraper.ArticleDom) entity.Article {

	var article entity.Article

	ArtInfo(dom.ProductInf, &article)
	SizeInfo(dom.SizeInf, &article)

	return article
}

func ArtInfo(dom string, article *entity.Article) {

	p := strings.NewReader(dom)
	var productInfo []string

	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("span, p").Each(func(i int, s *goquery.Selection) {
		info := s.Text()
		productInfo = append(productInfo, info)
		fmt.Printf("Pro info %d: %s\n", i, info)
	})
	article.SetProductInfo(productInfo)
}

func SizeInfo(dom string, article *entity.Article) {

	var productInfo []string
	p := strings.NewReader(dom)

	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div > span").Each(func(i int, s *goquery.Selection) {
		info := s.Text()
		productInfo = append(productInfo, info)
		fmt.Printf("size %d: %s\n", i, info)
	})
	article.SetSizeInfo(productInfo)
}
