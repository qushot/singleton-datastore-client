package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/qushot/singleton-datastore-client/domain/model"
	"github.com/qushot/singleton-datastore-client/domain/repository"
)

// NewUser returns constructor that implemented the repository.User interface.
func NewUser(connector Connector) repository.User {
	return &user{
		Connector: connector,
		kindName:  "User",
	}
}

type user struct {
	Connector
	kindName string
}

func (u user) Create(ctx context.Context, user *model.User) (*model.User, error) {
	client := u.Client()

	key := datastore.NameKey(u.kindName, user.Name, nil)
	if _, err := client.Put(ctx, key, user); err != nil {
		return nil, err
	}

	return user, nil
}
