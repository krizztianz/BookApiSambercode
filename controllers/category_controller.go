package controllers

import (
	"booklibraryapi/config"
	"booklibraryapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Categories
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Category
// @Router /api/categories [get]
func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)
		if err == nil {
			categories = append(categories, cat)
		}
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary Create Category
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.CategoryRequest true "Category data"
// @Success 201 {object} map[string]string
// @Router /api/categories [post]
func CreateCategory(c *gin.Context) {
	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := config.DB.Exec("INSERT INTO categories (name) VALUES ($1)", req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

// @Summary Get Books by Category ID
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {array} models.Book
// @Router /api/categories/{id}/books [get]
func GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")
	rows, err := config.DB.Query("SELECT id, title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by, modified_at, modified_by FROM books WHERE category_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.CategoryID, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
		if err == nil {
			books = append(books, b)
		}
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Get Category Detail
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /api/categories/{id} [get]
func GetCategoryDetail(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	err := config.DB.QueryRow("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1", id).Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// @Summary Delete Category
// @Tags Categories
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
