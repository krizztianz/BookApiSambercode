package controllers

import (
	"booklibraryapi/config"
	"booklibraryapi/models"
	"booklibraryapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/users/login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", req.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		//fmt.Println(req.Password + " : " + user.Password)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
