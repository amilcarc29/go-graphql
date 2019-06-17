package entities

import "github.com/jinzhu/gorm"

// Resource defines a struct that represents a Resource
type Resource struct {
	gorm.Model

	ID          string   `json:"id"`
	Link        string   `json:"link"`
	Name        string   `json:"name"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
