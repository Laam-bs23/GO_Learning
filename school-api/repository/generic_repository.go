package repository

import (
	"gorm.io/gorm"
)

// GenericRepository defines the interface for generic repository operations
type GenericRepository[T any] interface {
	Create(entity *T) error
	GetAll() ([]T, error)
	GetByID(id uint) (*T, error)
	Update(entity *T) error
	Delete(id uint) error
}

// genericRepository implements GenericRepository for any type T
type genericRepository[T any] struct {
	db *gorm.DB
}

// NewGenericRepository creates a new generic repository for type T
func NewGenericRepository[T any](db *gorm.DB) GenericRepository[T] {
	return &genericRepository[T]{db: db}
}

// Create adds a new entity to the database
func (r *genericRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// GetAll retrieves all entities of type T
func (r *genericRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

// GetByID retrieves an entity by its ID
func (r *genericRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update modifies an existing entity
func (r *genericRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Delete removes an entity by its ID
func (r *genericRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
} 