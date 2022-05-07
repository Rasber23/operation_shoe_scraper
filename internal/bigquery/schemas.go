package bigquery

import (
	"cloud.google.com/go/bigquery"
	"fmt"
)

func getSchema(schemaName string) (bigquery.Schema, error) {

	switch schemaName {
	case "articles":
		return bigquery.Schema{
			{Name: "category", Type: bigquery.StringFieldType},
			{Name: "brand", Type: bigquery.StringFieldType},
			{Name: "articleNumber", Type: bigquery.StringFieldType},
			{Name: "price", Type: bigquery.StringFieldType},
			{Name: "outerMaterial", Type: bigquery.StringFieldType},
			{Name: "lining", Type: bigquery.StringFieldType},
			{Name: "sole", Type: bigquery.StringFieldType},
			{Name: "color", Type: bigquery.StringFieldType},
		}, nil
	case "brands":
		return bigquery.Schema{}, nil
	default:
		return nil, fmt.Errorf("no schema with %v exsists", schemaName)
	}
}
