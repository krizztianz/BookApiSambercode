package routes

import (
	"booklibraryapi/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/users/login", controllers.Login)
	r.GET("/users/authenticated", controllers.Authenticated)
}
