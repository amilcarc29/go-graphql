package roles

import (
	resourceHandler "go-graphql/src/api/domain/resource/delivery/http"
	"go-graphql/src/api/infraestructure/dependencies"
)

// Reader defines a reader struct
type Reader struct {
	container *dependencies.Container
}

// NewReader returns a new reader
func NewReader(container *dependencies.Container) *Reader {
	return &Reader{
		container: container,
	}
}

// RegisterRoutes registers read routes
func (reader *Reader) RegisterRoutes(basePath string) {
	resourceHandler := resourceHandler.NewResourceHandler(reader.container)

	routerHandler := reader.container.RouterHandler()

	routerHandler.HandleFunc("/resources", resourceHandler.GetResources).Methods("GET")
	routerHandler.HandleFunc("/resources/{id}", resourceHandler.GetResource).Methods("GET")
}
