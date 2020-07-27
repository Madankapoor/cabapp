package apiv1

import (
	"github.com/Madankapoor/cabapp/backendapi/tripservice/api/v1.0/trip"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		trip.ApplyRoutes(v1)
	}
}
