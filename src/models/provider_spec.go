package models

import (
	"cloud.google.com/go/datastore"
	"time"
)

type ProviderSpec struct {
	K            *datastore.Key `json:"id,omitempty" datastore:"__key__" `
	Region       string         `json:"region" :"name"`
	Zone         string         `json:"zone" :"path"`
	ProviderName string         `json:"provider_name" :"size"`
	ProviderUrl  string         `json:"provider_url" :"hash"`
	CreatedAt    time.Time      `json:"created_at,omitempty" :"created_at"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty" :"updated_at"`
}
