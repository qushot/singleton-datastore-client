package registry

import "github.com/qushot/singleton-datastore-client/domain/repository"

// Registry represents DI Container.
type Registry interface {
	// Repositories
	GetUserRepository() repository.User
}

var reg Registry

// Get returns DI Container that implements registry.Registry interface.
func Get() (Registry, error) {
	if reg == nil {
		r, err := initialize()
		if err != nil {
			return nil, err
		}
		reg = r
	}
	return reg, nil
}

// registry is the DI Container.
type registry struct {
	userRepository repository.User
}

func newRegistry(userRepository repository.User) Registry {
	return &registry{userRepository: userRepository}
}

func (r *registry) GetUserRepository() repository.User {
	return r.userRepository
}
