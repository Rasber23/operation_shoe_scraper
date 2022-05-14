package entity

import (
	"cloud.google.com/go/bigquery"
	"fmt"
	"strconv"
	"strings"
)

type Article struct {
	Brand         string
	Category      string
	ArticleNumber string
	ArticleName   string
	Price         float64
	OuterMaterial string
	Lining        string
	Sole          string
	SizeInfo      string
}

type Articles []Article

func (article Article) Save() map[string]bigquery.Value {
	return map[string]bigquery.Value{
		"Brand":         article.Brand,
		"Category":      article.Category,
		"ArticleNumber": article.ArticleNumber,
		"ArticleName":   article.ArticleName,
		"Price":         article.Price,
		"OuterMaterial": article.OuterMaterial,
		"Lining":        article.Lining,
		"Sole":          article.Sole,
		"SizeInfo":      article.SizeInfo,
	}
}

func (article *Article) SetProductInfo(info []string) {

	article.Brand = info[0]
	article.Category = "not Implemented"
	article.ArticleNumber = info[44]
	article.ArticleName = info[1]
	article.Price = stringToFloat(info[4])
	article.OuterMaterial = info[27]
	article.Lining = info[29]
	article.Sole = info[31]

}
func (article *Article) SetSizeInfo(info []string) {
	article.SizeInfo = strings.Join(info, ",")
}
func stringToFloat(price string) float64 {

	price = price[:len(price)-2]
	price = strings.ReplaceAll(price, "\u00a0", "")
	price = strings.ReplaceAll(price, ",", ".")

	result, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Errorf("could not transform %v", price)
	}
	return result
}
