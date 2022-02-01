package datasources

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"incorta-cloud/upload-service/config"
	"os"
)

var client *datastore.Client

func InitDatastore() {
	fmt.Println("Initializing Datastore connection.")
	ctx := context.Background()
	// Set up admin client, tables, and column families.
	// NewAdminClient uses Application Default Credentials to authenticate.
	val := config.GetConfig().GetString("DATASTORE_PROJECT_ID")

	if val == "" {
		fmt.Println("Missing DATASTORE_PROJECT_ID")
		os.Exit(1)
	}
	fmt.Println("Connecting To Project " + val)

	clnt, err := datastore.NewClient(ctx, val)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client = clnt

	//oauthHttpClient, err = google.DefaultClient(ctx, "https://www.googleapis.com/auth/devstorage.read_write")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	fmt.Println("Initialization of Datastore Complete.")

}

func GetClient() *datastore.Client {
	return client
}
