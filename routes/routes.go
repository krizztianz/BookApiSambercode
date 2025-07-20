package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	UserRoutes(api)
	BookRoutes(api)
	CategoryRoutes(api)
	// Add category route here
}
