package http

import (
	"go-graphql/src/api/infraestructure/dependencies"

	"github.com/gorilla/mux"
)

// ResourceHTTPRepository defines a resource http repository struct
type ResourceHTTPRepository struct {
	http *mux.Router
}

// NewRepository returns a new resource http repository
func NewRepository(container *dependencies.Container) *ResourceHTTPRepository {

	return &ResourceHTTPRepository{
		http: container.RouterHandler(),
	}
}
