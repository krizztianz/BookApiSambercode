package routes

import (
	"booklibraryapi/controllers"
	"booklibraryapi/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.RouterGroup) {
	book := r.Group("/books")
	book.Use(middleware.JWTAuthMiddleware())
	{
		book.GET("", controllers.GetBooks)
		book.POST("", controllers.CreateBook)
		book.GET(":id", controllers.GetBookDetail)
		book.DELETE(":id", controllers.DeleteBook)
	}
}
