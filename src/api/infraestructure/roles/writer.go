package roles

import (
	resourceHandler "go-graphql/src/api/domain/resource/delivery/http"
	"go-graphql/src/api/infraestructure/dependencies"
)

// Writer defines a writer struct
type Writer struct {
	container *dependencies.Container
}

// NewWriter returns a new writer
func NewWriter(container *dependencies.Container) *Writer {
	return &Writer{
		container: container,
	}
}

// RegisterRoutes registers write routes
func (writer *Writer) RegisterRoutes(basePath string) {
	resourceHandler := resourceHandler.NewResourceHandler(writer.container)

	routerHandler := writer.container.RouterHandler()

	routerHandler.HandleFunc("/resources", resourceHandler.CreateResource).Methods("POST")
	routerHandler.HandleFunc("/resources/{id}", resourceHandler.DeleteResource).Methods("DELETE")
}
