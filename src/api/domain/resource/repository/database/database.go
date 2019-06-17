package database

import (
	"go-graphql/src/api/infraestructure/dependencies"

	"github.com/jinzhu/gorm"
)

// ResourceDatabaseRepository defines a repository struct
type ResourceDatabaseRepository struct {
	database *gorm.DB
}

// NewRepository returns a new Resource repository
func NewRepository(container *dependencies.Container) *ResourceDatabaseRepository {
	return &ResourceDatabaseRepository{
		database: container.Database(),
	}
}

// DatabaseBegin begins a transaction
func (repository *ResourceDatabaseRepository) DatabaseBegin() *gorm.DB {
	return repository.database.Begin()
}

// DatabaseCommit commits a transaction
func (repository *ResourceDatabaseRepository) DatabaseCommit() *gorm.DB {
	return repository.database.Commit()
}

// DatabaseRollback rollbacks a transaction
func (repository *ResourceDatabaseRepository) DatabaseRollback() *gorm.DB {
	return repository.database.Rollback()
}

// GetNewUUID returns a new UUID
func (repository *ResourceDatabaseRepository) GetNewUUID() string {
	var newID string
	repository.database.DB().QueryRow("SELECT UUID()", &newID)
	return newID
}
