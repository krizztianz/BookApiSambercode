package controllers

import (
	"booklibraryapi/config"
	"booklibraryapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Books
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Book
// @Router /api/books [get]
func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by, modified_at, modified_by FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books", "details": err.Error()})
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

// @Summary Create Book
// @Tags Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.BookRequest true "Book data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/books [post]
func CreateBook(c *gin.Context) {
	var req models.BookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	thickness := "Tipis"
	if req.TotalPage >= 100 {
		thickness = "Tebal"
	}

	_, err := config.DB.Exec("INSERT INTO books (title, category_id, description, image_url, release_year, price, total_page, thickness) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)",
		req.Title, req.CategoryID, req.Description, req.ImageURL, req.ReleaseYear, req.Price, req.TotalPage, thickness)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

// @Summary Get Book Detail
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [get]
func GetBookDetail(c *gin.Context) {
	id := c.Param("id")
	var b models.Book
	err := config.DB.QueryRow("SELECT id, title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by, modified_at, modified_by FROM books WHERE id=$1", id).Scan(&b.ID, &b.Title, &b.CategoryID, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

// @Summary Delete Book
// @Tags Books
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
