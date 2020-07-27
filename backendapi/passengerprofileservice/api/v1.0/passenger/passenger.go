package passenger

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	passenger := r.Group("/passenger")
	{
		passenger.POST("/", create)
		passenger.GET("/", list)
		passenger.GET("/:id", read)
		passenger.DELETE("/:id", remove)
		passenger.PATCH("/:id", update)
	}
}
