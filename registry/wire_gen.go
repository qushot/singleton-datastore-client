// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/qushot/singleton-datastore-client/infrastructure/persistence/datastore"
)

// Injectors from wire.go:

func initialize() (Registry, error) {
	connector, err := datastore.NewConnector()
	if err != nil {
		return nil, err
	}
	user := datastore.NewUser(connector)
	registryRegistry := newRegistry(user)
	return registryRegistry, nil
}

// wire.go:

var repositorySet = wire.NewSet(datastore.NewUser, datastore.NewConnector)
