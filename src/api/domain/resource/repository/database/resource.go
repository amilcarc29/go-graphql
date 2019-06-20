package database

import (
	"fmt"
	"go-graphql/src/api/domain/resource/entities"
	"strconv"
)

const (
	errResourceNotFound = "Resource %s not found"
)

// GetResources returns all the resources found
func (repository *ResourceDatabaseRepository) GetResources() ([]entities.Resource, error) {
	var resources []entities.Resource
	if err := repository.database.Find(&resources).Error; err != nil {
		return nil, err
	}
	return resources, nil
}

// GetResource returns the resource found
func (repository *ResourceDatabaseRepository) GetResource(id string) (entities.Resource, error) {
	var resource entities.Resource
	idConverted, _ := strconv.Atoi(id)
	err := repository.database.First(&resource, idConverted).Error
	if resource.ID == 0 || err != nil {
		return resource, fmt.Errorf(errResourceNotFound, id)
	}
	return resource, nil
}

// PutResource inserts a new resource
func (repository *ResourceDatabaseRepository) PutResource(resource entities.Resource) (uint, error) {
	if err := repository.database.Create(&resource).Error; err != nil {
		return 0, err
	}
	return resource.ID, nil
}

// DeleteResource deletes a resource
func (repository *ResourceDatabaseRepository) DeleteResource(id string) error {
	resource, _ := repository.GetResource(id)
	if err := repository.database.Delete(&resource).Error; err != nil {
		return err
	}
	return nil
}
