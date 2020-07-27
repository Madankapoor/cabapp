package cabdriver

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	cabdriver := r.Group("/cabdriver")
	{
		cabdriver.POST("/", create)
		cabdriver.GET("/", list)
		cabdriver.GET("/:id", read)
		cabdriver.DELETE("/:id", remove)
		cabdriver.PATCH("/:id", update)
	}
}
