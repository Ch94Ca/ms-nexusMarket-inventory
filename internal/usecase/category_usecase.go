// Package usecase implements the business logic for category management.
package usecase

import (
	"context"
	"time"

	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/app/dtos"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/domain"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) CreateCategory(ctx context.Context, dto dtos.CreateCategoryDTO) (*domain.Category, error) {
	if dto.Name == "" {
		return nil, domain.ErrInvalidCategoryName
	}

	category := &domain.Category{
		Name:      dto.Name,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (u *CategoryUsecase) ListCategories(ctx context.Context) ([]*domain.Category, error) {
	return u.repo.ListAll()
}

func (u *CategoryUsecase) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	return u.repo.GetByID(id)
}

func (u *CategoryUsecase) UpdateCategory(ctx context.Context, id int, dto dtos.UpdateCategoryDTO) error {

	if dto.Name == "" {
		return domain.ErrInvalidCategoryName
	}

	category := &domain.Category{
		ID:   id,
		Name: dto.Name,
	}

	return u.repo.Update(category)
}

func (u *CategoryUsecase) DeleteCategory(ctx context.Context, id int) error {
	return u.repo.Delete(id)
}
