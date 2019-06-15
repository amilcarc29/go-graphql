package database

import (
	"go-graphql/src/api/domain/resource/entities"
)

// GetResources TODO
func (repository *ResourceDatabaseRepository) GetResources() ([]entities.Resource, error) {
	return []entities.Resource{}, nil
}

// GetResource TODO
func (repository *ResourceDatabaseRepository) GetResource(id uint64) (entities.Resource, error) {
	return entities.Resource{}, nil
}

// PutResource TODO
func (repository *ResourceDatabaseRepository) PutResource(entities.Resource) error {
	return nil
}

// DeleteResource TODO
func (repository *ResourceDatabaseRepository) DeleteResource(id uint64) error {
	return nil
}
