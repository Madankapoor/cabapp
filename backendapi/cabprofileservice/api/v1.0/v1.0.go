package apiv1

import (
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/api/v1.0/cab"
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/api/v1.0/cabdriver"
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/api/v1.0/cabtype"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/health", health)
		cab.ApplyRoutes(v1)
		cabdriver.ApplyRoutes(v1)
		cabtype.ApplyRoutes(v1)
	}
}
