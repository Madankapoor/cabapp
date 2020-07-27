package apiv1

import (
	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/api/v1.0/cablocation"
	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/api/v1.0/passengerlocation"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/health", healthCheck)
		cablocation.ApplyRoutes(v1)
		passengerlocation.ApplyRoutes(v1)
	}
}
