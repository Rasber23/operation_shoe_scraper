package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
)

func getTable() {

}

func createTable(projectId, datasetID, tableID string, schema bigquery.Schema) error {

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	metaData := &bigquery.TableMetadata{
		Schema: schema,
	}

	tableRef := client.Dataset(datasetID).Table(tableID)
	if err := tableRef.Create(ctx, metaData); err != nil {
		return err
	}
	return nil
}
