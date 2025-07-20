package routes

import (
	"booklibraryapi/controllers"
	"booklibraryapi/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	category := r.Group("/categories")
	category.Use(middleware.JWTAuthMiddleware())
	{
		category.GET("", controllers.GetCategories)
		category.POST("", controllers.CreateCategory)
		category.GET(":id", controllers.GetCategoryDetail)
		category.GET(":id/books", controllers.GetBooksByCategory)
		category.DELETE(":id", controllers.DeleteCategory)
	}
}
