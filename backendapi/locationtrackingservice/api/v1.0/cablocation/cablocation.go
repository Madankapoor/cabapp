package cablocation

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	cablocation := r.Group("/cablocation")
	{
		cablocation.POST("/", create)
		cablocation.GET("/", list)
		cablocation.GET("/:id", read)
		cablocation.DELETE("/:id", remove)
		cablocation.PATCH("/:id", update)
	}
}
