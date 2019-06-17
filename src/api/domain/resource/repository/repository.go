package repository

import (
	"go-graphql/src/api/domain/resource/entities"
	database "go-graphql/src/api/domain/resource/repository/database"
	"go-graphql/src/api/domain/resource/repository/http"
	"go-graphql/src/api/infraestructure/dependencies"
)

// ResourceRepository defines an interface
type ResourceRepository interface {
	GetResources() ([]entities.Resource, error)
	GetResource(id string) (entities.Resource, error)
	PutResource(entities.Resource) error
	DeleteResource(id string) error
}

// Repository defines a repository struct
type Repository struct {
	*database.ResourceDatabaseRepository
	*http.ResourceHTTPRepository
}

// NewResourceRepository returns a new resource repository
func NewResourceRepository(container *dependencies.Container) ResourceRepository {
	return &Repository{
		database.NewRepository(container),
		http.NewRepository(container),
	}
}
