package datasources

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"os"
)

var storageClient *storage.Client

func InitStorage() {
	ctx := context.Background()
	// Set up admin client, tables, and column families.
	// NewAdminClient uses Application Default Credentials to authenticate.
	//adminClient, err := bigtable.NewAdminClient(ctx, "", "")
	clnt, err := storage.NewClient(ctx)

	fmt.Println(clnt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	storageClient = clnt

}

func GetStorageClient() *storage.Client {
	return storageClient
}
