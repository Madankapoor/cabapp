package cabtype

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	cabtype := r.Group("/cabtype")
	{
		cabtype.POST("/", create)
		cabtype.GET("/", list)
		cabtype.GET("/:id", read)
		cabtype.DELETE("/:id", remove)
		cabtype.PATCH("/:id", update)
	}
}
