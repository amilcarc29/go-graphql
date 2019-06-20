package main

import (
	"fmt"
	resourceHandler "go-graphql/src/api/domain/resource/delivery/http"
	"go-graphql/src/api/infraestructure/dependencies"
	"log"
	"net/http"
)

func main() {
	container, err := dependencies.NewContainer()

	if err != nil {
		fmt.Printf(err.Error())
	}

	resourceHandler := resourceHandler.NewResourceHandler(container)

	routerHandler := container.RouterHandler()
	routerHandler.HandleFunc("/resources", resourceHandler.CreateResource).Methods("POST")
	routerHandler.HandleFunc("/resources/{id}", resourceHandler.DeleteResource).Methods("DELETE")
	routerHandler.HandleFunc("/resources", resourceHandler.GetResources).Methods("GET")
	routerHandler.HandleFunc("/resources/{id}", resourceHandler.GetResource).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", routerHandler))
}
