package trip

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	trip := r.Group("/trip")
	{
		trip.POST("/", createtrip)
		trip.GET("/", listtrips)
		trip.GET("/:passenger_id", readtrip)
		trip.PATCH("/:id", update)
	}
}
