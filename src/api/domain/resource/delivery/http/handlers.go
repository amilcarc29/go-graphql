package http

import (
	"go-graphql/src/api/domain/resource/usecases"
	"go-graphql/src/api/infraestructure/dependencies"
)

// ResourceHandler defines a resource handler struct
type ResourceHandler struct {
	usecases *usecases.UseCases
}

// NewResourceHandler returns a new resource handler
func NewResourceHandler(container *dependencies.Container) *ResourceHandler {
	usecases := usecases.NewUseCases(container)
	return &ResourceHandler{
		usecases: usecases,
	}
}
