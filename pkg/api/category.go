package api

import (
	"owwi/pkg/models"
	"owwi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param register body models.CreateCategory true "Category data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /categories [post]
func createCategory(c *gin.Context) {
	var categoryData models.CreateCategory
	if err := c.BindJSON(&categoryData); err != nil || categoryData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := repositories.CategoryRepository.CreateCategory(categoryData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create type"})
		return
	}

	c.Status(201)
}

// @Summary Update Category
// @Description Update Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param register body models.UpdateCategory true "Category data"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /categories [put]
func updateCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Category ID is required"})
		return
	}
	var categoryData models.UpdateCategory
	if err := c.BindJSON(&categoryData); err != nil || categoryData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	categoryData.ID = id
	if err := repositories.CategoryRepository.UpdateCategory(categoryData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.Status(200)
}

// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /categories/:id [get]
func getCategoryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Category ID is required"})
		return
	}

	category, err := repositories.CategoryRepository.GetCategoryByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve category", "err": err.Error()})
		return
	}

	if category == nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, category)
}

// @Summary Get All Categories By User
// @Description Get all categories associated with the authenticated user
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /categories [get]
func getAllCategoriesByUser(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	categories, err := repositories.CategoryRepository.GetAllCategoriesByUser(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	if len(categories) == 0 {
		c.JSON(404, gin.H{"message": "No categories found for this user", "user_id": userID, "categories": categories})
		return
	}

	c.JSON(200, categories)
}
// @Summary Delete Category
// @Description Delete Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Security BearerAuth
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /categories/:id [delete]
func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Category ID is required"})
		return
	}

	if err := repositories.CategoryRepository.DeleteCategory(id); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete category"})
		return
	}

	c.Status(204)
}

var CategoryApi = struct {
	CreateCategory         func(*gin.Context)
	UpdateCategory         func(*gin.Context)
	GetCategoryByID        func(*gin.Context)
	GetAllCategoriesByUser func(*gin.Context)
	DeleteCategory         func(*gin.Context)
}{
	CreateCategory:         createCategory,
	UpdateCategory:         updateCategory,
	GetCategoryByID:        getCategoryByID,
	GetAllCategoriesByUser: getAllCategoriesByUser,
	DeleteCategory:         deleteCategory,
}
