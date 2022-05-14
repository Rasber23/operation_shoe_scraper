package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/Rasber23/operation_shoe_scraper/internal/business/entity"
)

type BQService struct {
	ProjectId string
	DataSetId string
	TableId   string
}

func (service BQService) SaveArticles(articles entity.Articles) {

	/*	err := CreateDataset(service.ProjectId, service.DataSetId)
		if err != nil {
			fmt.Printf("Error Creating Dataset: %v", err)
		}

		schema, err := getSchema("articles")
		if err != nil {
			fmt.Printf("Error Creating Schema: %v", err)
		}

		err = createTable(service.ProjectId, service.DataSetId, service.ProjectId, schema)
		if err != nil {
			fmt.Printf("Error Creating Table: %v", err)
		}*/

	err := insertRows(service.ProjectId, service.DataSetId, service.ProjectId, articles)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
}

func insertRows(projectID, datasetID, tableID string, articles entity.Articles) error {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	inserter := client.Dataset(datasetID).Table(tableID).Inserter()
	resultArticles := entity.Articles{}
	for _, article := range articles {
		resultArticles = append(resultArticles, article)
	}

	if err := inserter.Put(ctx, resultArticles); err != nil {
		return err
	}
	return nil
}
