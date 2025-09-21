// Package dtos defines Data Transfer Objects used in the application.
package dtos

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

type CategoryDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
