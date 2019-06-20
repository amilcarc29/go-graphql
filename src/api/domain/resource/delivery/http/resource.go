package http

import (
	"encoding/json"
	"go-graphql/src/api/domain/resource/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *ResourceHandler) GetResources(w http.ResponseWriter, r *http.Request) {

	resources, err := handler.usecases.GetResources()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&resources)
	w.WriteHeader(http.StatusOK)
}

func (handler *ResourceHandler) GetResource(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	resource, err := handler.usecases.GetResource(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&resource)
	w.WriteHeader(http.StatusOK)
}

func (handler *ResourceHandler) CreateResource(w http.ResponseWriter, r *http.Request) {

	var resource entities.Resource
	json.NewDecoder(r.Body).Decode(&resource)
	newID, err := handler.usecases.CreateResource(resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resourceIDResponse := struct {
		ID uint `json:"id"`
	}{
		ID: newID,
	}

	json.NewEncoder(w).Encode(&resourceIDResponse)
	w.WriteHeader(http.StatusCreated)
}

func (handler *ResourceHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	err := handler.usecases.DeleteResource(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
