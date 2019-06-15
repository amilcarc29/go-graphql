package usecases

import (
	"go-graphql/src/api/domain/resource/repository"
	"go-graphql/src/api/infraestructure/dependencies"
)

// UseCases defines a usecases struct
type UseCases struct {
	resourceRepository repository.ResourceRepository
}

// NewUseCases returns a new usecases
func NewUseCases(container *dependencies.Container) *UseCases {
	resourceRepository := repository.NewResourceRepository(container)
	return &UseCases{
		resourceRepository: resourceRepository,
	}
}
