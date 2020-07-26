package apiv1

import (
	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/api/v1.0/passenger"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		passenger.ApplyRoutes(v1)
	}
}
