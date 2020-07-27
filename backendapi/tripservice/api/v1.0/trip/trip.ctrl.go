package trip

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/tripservice/models"
	tripservice "github.com/Madankapoor/cabapp/backendapi/tripservice/service/trip"
	"github.com/gin-gonic/gin"
)

func createtrip(c *gin.Context) {
	var tripRequest models.TripRequest
	if err := c.BindJSON(&tripRequest); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": false, "message": err.Error(),
		})
		return
	}
	trip, err := tripservice.BookTrip(c, tripRequest.GetTripModel())
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": false, "message": err.Error(),
		})
		return
	}
	c.JSON(200, *trip)
}

func listtrips(c *gin.Context) {
	trips, err := tripservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, trips)
}

func readtrip(c *gin.Context) {
	strid := c.Param("passenger_id")
	id, _ := strconv.Atoi(strid)
	trip, err := tripservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, trip)
}

func update(c *gin.Context) {
	var trip models.Trip
	if err := c.BindJSON(&trip); err != nil {
		c.AbortWithStatus(400)
		return
	}
	trip, err := tripservice.Update(c, &trip)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, trip)
}
