package bigquery

import (
	"cloud.google.com/go/bigquery"
	"fmt"
)

func getSchema(schemaName string) (bigquery.Schema, error) {

	switch schemaName {
	case "articles":
		return bigquery.Schema{
			{Name: "Brand", Type: bigquery.StringFieldType},
			{Name: "Category", Type: bigquery.StringFieldType},
			{Name: "ArticleNumber", Type: bigquery.StringFieldType},
			{Name: "ArticleName", Type: bigquery.StringFieldType},
			{Name: "Price", Type: bigquery.FloatFieldType},
			{Name: "OuterMaterial", Type: bigquery.StringFieldType},
			{Name: "Lining", Type: bigquery.StringFieldType},
			{Name: "Sole", Type: bigquery.StringFieldType},
			{Name: "SizeInfo", Type: bigquery.StringFieldType},
		}, nil
	default:
		return nil, fmt.Errorf("no schema with %v exsists", schemaName)
	}
}
