package usecases

import "go-graphql/src/api/domain/resource/entities"

// GetResources returns the resources
func (usecases *UseCases) GetResources() ([]entities.Resource, error) {
	resources, err := usecases.resourceRepository.GetResources()
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// GetResource returns the resource
func (usecases *UseCases) GetResource(id string) (*entities.Resource, error) {
	resource, err := usecases.resourceRepository.GetResource(id)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// CreateResource creates a new resource
func (usecases *UseCases) CreateResource(resource entities.Resource) (uint, error) {
	id, err := usecases.resourceRepository.PutResource(resource)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// DeleteResource deletes the resource
func (usecases *UseCases) DeleteResource(id string) error {
	err := usecases.resourceRepository.DeleteResource(id)
	if err != nil {
		return err
	}
	return nil
}
