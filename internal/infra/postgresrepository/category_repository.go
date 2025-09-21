// package postgresrepository implements the category repository using PostgreSQL and GORM.
package postgresrepository

import (
	"errors"

	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryPostgres struct {
	db *gorm.DB
}

func NewCategoryRepositoryPostgres(db *gorm.DB) *CategoryRepositoryPostgres {
	return &CategoryRepositoryPostgres{
		db: db,
	}
}

func (r *CategoryRepositoryPostgres) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepositoryPostgres) GetByID(id int) (*domain.Category, error) {
	var category domain.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, result.Error
	}
	return &category, nil
}

func (r *CategoryRepositoryPostgres) ListAll() ([]*domain.Category, error) {
	var categories []*domain.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *CategoryRepositoryPostgres) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepositoryPostgres) Delete(id int) error {
	return r.db.Delete(&domain.Category{}, id).Error
}
