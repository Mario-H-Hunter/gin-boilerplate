package seeders

import "cloud.google.com/go/datastore"

type DataStoreSeeder interface {
	seed(client datastore.Client)
}
