package passengerlocation

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	passengerlocation := r.Group("/passengerlocation")
	{
		passengerlocation.POST("/", create)
		passengerlocation.GET("/", list)
		passengerlocation.GET("/:id", read)
		passengerlocation.DELETE("/:id", remove)
		passengerlocation.PATCH("/:id", update)
	}
}
