//+build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/qushot/singleton-datastore-client/infrastructure/persistence/datastore"
)

var repositorySet = wire.NewSet(
	datastore.NewUser,
	datastore.NewConnector,
)

func initialize() (Registry, error) {
	wire.Build(
		newRegistry,
		repositorySet,
	)
	return &registry{}, nil
}
