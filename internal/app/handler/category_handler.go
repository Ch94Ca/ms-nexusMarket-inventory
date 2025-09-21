// Package handler implements HTTP handlers for the application.
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/app/dtos"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/domain"
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase *usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

// CreateCategory creates a new category
// @Summary Create Category
// @Description Creates a new category
// @ID create_category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dtos.CreateCategoryDTO true "Category to be created"
// @Success 201 {object} dtos.CategoryDTO
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var createCategoryDTO dtos.CreateCategoryDTO

	if err := c.ShouldBindJSON(&createCategoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	category, err := h.categoryUsecase.CreateCategory(c.Request.Context(), createCategoryDTO)
	if err != nil {
		switch err {
		case domain.ErrInvalidCategoryName:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, category)
}

// ListCategories lists all categories
// @Summary List Categories
// @Description Retrieves a list of all categories
// @ID list_categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} dtos.CategoryDTO "List of categories"
// @Router /categories [get]
func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.categoryUsecase.ListCategories(c.Request.Context())
	if err != nil {
		switch err {
		case domain.ErrInvalidCategoryName:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID find category by ID
// @Summary Find Category by ID
// @Description Retrieves a category by its ID
// @ID find_category_by_id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dtos.CategoryDTO "Category found"
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	category, err := h.categoryUsecase.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error retrieving category: %v", err)
		switch err {
		case domain.ErrCategoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category
// @Summary Update Category
// @Description Updates an existing category
// @ID update_category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dtos.UpdateCategoryDTO true "Category to be updated"
// @Success 204 {object} dtos.CategoryDTO
// @Router /categories/{id} [patch]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateCategoryDTO dtos.UpdateCategoryDTO

	if err := c.ShouldBindJSON(&updateCategoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.categoryUsecase.UpdateCategory(c.Request.Context(), id, updateCategoryDTO)
	if err != nil {
		switch err {
		case domain.ErrInvalidCategoryName:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteCategory deletes an existing category
// @Summary Delete Category
// @Description Deletes an existing category
// @ID delete_category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 {object} dtos.CategoryDTO
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.categoryUsecase.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error deleting category: %v", err)
		switch err {
		case domain.ErrCategoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
