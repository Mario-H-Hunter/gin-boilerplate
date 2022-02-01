package utils

import (
	"cloud.google.com/go/datastore"
	"incorta-cloud/upload-service/config"
)

func appNamespace() string {
	return config.GetConfig().GetString("server.datastore-namespace")
}

func Query(kind string) *datastore.Query {
	q := datastore.NewQuery(kind)
	q = q.Namespace(appNamespace())
	return q
}

func DecodeKey(encodedId string) *datastore.Key {
	k, err := datastore.DecodeKey(encodedId)
	if err != nil {
		return nil

	}
	k.Namespace = appNamespace()
	return k
}

func IdKey(kind string, id int64, parent *datastore.Key) *datastore.Key {
	k := datastore.IDKey(kind, id, parent)
	k.Namespace = appNamespace()
	return k
}

func NameKey(kind string, name string, parent *datastore.Key) *datastore.Key {
	k := datastore.NameKey(kind, name, parent)
	k.Namespace = appNamespace()
	return k
}

func IncompleteKey(kind string, parent *datastore.Key) *datastore.Key {
	k := datastore.IncompleteKey(kind, parent)
	k.Namespace = appNamespace()
	return k
}
