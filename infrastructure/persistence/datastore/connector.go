package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Connector represents a connector of cloud datastore.
type Connector interface {
	Client() *datastore.Client
}

// NewConnector returns constructor that implemented the datastore.Connector interface.
func NewConnector() (Connector, error) {
	client, err := datastore.NewClient(context.Background(), datastore.DetectProjectID)
	if err != nil {
		return nil, err
	}

	return &connector{
		client: client,
	}, nil
}

type connector struct {
	client *datastore.Client
}

func (c *connector) Client() *datastore.Client {
	return c.client
}
