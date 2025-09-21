package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/app/dtos"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/domain"
	categoryRepositoryMock "github.com/Ch94Ca/ms-nexusMarket-inventory/internal/mocks/domain/CategoryRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategory(t *testing.T) {
	tests := []struct {
		name         string
		dto          dtos.CreateCategoryDTO
		mockSetup    func(repo *categoryRepositoryMock.MockCategoryRepository)
		expectedErr  error
		expectedName string
	}{
		{
			name: "success",
			dto:  dtos.CreateCategoryDTO{Name: "Electronics"},
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("Create", mock.MatchedBy(func(cat *domain.Category) bool {
					return cat.Name == "Electronics"
				})).Return(nil)
			},
			expectedErr:  nil,
			expectedName: "Electronics",
		},
		{
			name:        "invalid name",
			dto:         dtos.CreateCategoryDTO{Name: ""},
			mockSetup:   func(repo *categoryRepositoryMock.MockCategoryRepository) {}, // Não espera chamada
			expectedErr: domain.ErrInvalidCategoryName,
		},
		{
			name: "repo error",
			dto:  dtos.CreateCategoryDTO{Name: "Electronics"},
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("Create", mock.Anything).Return(errors.New("erro do banco"))
			},
			expectedErr: errors.New("erro do banco"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := categoryRepositoryMock.NewMockCategoryRepository(t)
			if tc.mockSetup != nil {
				tc.mockSetup(repo)
			}
			usecase := NewCategoryUsecase(repo)
			ctx := context.Background()
			category, err := usecase.CreateCategory(ctx, tc.dto)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				if errors.Is(tc.expectedErr, domain.ErrInvalidCategoryName) {
					assert.ErrorIs(t, err, domain.ErrInvalidCategoryName)
				}
				assert.Nil(t, category)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, category)
				assert.Equal(t, tc.expectedName, category.Name)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestListCategories(t *testing.T) {
	tests := []struct {
		name         string
		mockSetup    func(repo *categoryRepositoryMock.MockCategoryRepository)
		expectedErr  error
		expectedList []*domain.Category
	}{
		{
			name: "success",
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("ListAll").Return([]*domain.Category{
					{ID: 1, Name: "Electronics"},
					{ID: 2, Name: "Books"},
				}, nil)
			},
			expectedErr: nil,
			expectedList: []*domain.Category{
				{ID: 1, Name: "Electronics"},
				{ID: 2, Name: "Books"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := categoryRepositoryMock.NewMockCategoryRepository(t)
			if tc.mockSetup != nil {
				tc.mockSetup(repo)
			}
			usecase := NewCategoryUsecase(repo)
			ctx := context.Background()
			categories, err := usecase.ListCategories(ctx)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, categories)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedList, categories)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestGetCategoryByID(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		mockSetup   func(repo *categoryRepositoryMock.MockCategoryRepository)
		expectedErr error
		expectedCat *domain.Category
	}{
		{
			name: "success",
			id:   1,
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("GetByID", 1).Return(&domain.Category{ID: 1, Name: "Electronics"}, nil)
			},
			expectedErr: nil,
			expectedCat: &domain.Category{ID: 1, Name: "Electronics"},
		},
		{
			name: "not found",
			id:   2,
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("GetByID", 2).Return(nil, domain.ErrCategoryNotFound)
			},
			expectedErr: domain.ErrCategoryNotFound,
			expectedCat: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := categoryRepositoryMock.NewMockCategoryRepository(t)
			if tc.mockSetup != nil {
				tc.mockSetup(repo)
			}
			usecase := NewCategoryUsecase(repo)
			ctx := context.Background()
			category, err := usecase.GetCategoryByID(ctx, tc.id)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, category)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, category)
				assert.Equal(t, tc.expectedCat, category)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		dto         dtos.UpdateCategoryDTO
		mockSetup   func(repo *categoryRepositoryMock.MockCategoryRepository)
		expectedErr error
	}{
		{
			name: "success",
			id:   1,
			dto:  dtos.UpdateCategoryDTO{Name: "Updated Electronics"},
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("Update", mock.MatchedBy(func(cat *domain.Category) bool {
					return cat.ID == 1 && cat.Name == "Updated Electronics"
				})).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:        "invalid name",
			id:          1,
			dto:         dtos.UpdateCategoryDTO{Name: ""},
			mockSetup:   func(repo *categoryRepositoryMock.MockCategoryRepository) {}, // Não espera chamada
			expectedErr: domain.ErrInvalidCategoryName,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := categoryRepositoryMock.NewMockCategoryRepository(t)
			if tc.mockSetup != nil {
				tc.mockSetup(repo)
			}
			usecase := NewCategoryUsecase(repo)
			ctx := context.Background()
			err := usecase.UpdateCategory(ctx, tc.id, tc.dto)
			if tc.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		mockSetup   func(repo *categoryRepositoryMock.MockCategoryRepository)
		expectedErr error
	}{
		{
			name: "success",
			id:   1,
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("Delete", 1).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			id:   2,
			mockSetup: func(repo *categoryRepositoryMock.MockCategoryRepository) {
				repo.On("Delete", 2).Return(domain.ErrCategoryNotFound)
			},
			expectedErr: domain.ErrCategoryNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := categoryRepositoryMock.NewMockCategoryRepository(t)
			if tc.mockSetup != nil {
				tc.mockSetup(repo)
			}
			usecase := NewCategoryUsecase(repo)
			ctx := context.Background()
			err := usecase.DeleteCategory(ctx, tc.id)
			if tc.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			repo.AssertExpectations(t)
		})
	}
}
