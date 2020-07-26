package passenger

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.POST("/", create)
		posts.GET("/", list)
		posts.GET("/:id", read)
		posts.DELETE("/:id", remove)
		posts.PATCH("/:id", update)
	}
}
