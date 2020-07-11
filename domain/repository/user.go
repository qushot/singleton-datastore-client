package repository

import (
	"context"

	"github.com/qushot/singleton-datastore-client/domain/model"
)

// User represents a repository of an user.
type User interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
}
