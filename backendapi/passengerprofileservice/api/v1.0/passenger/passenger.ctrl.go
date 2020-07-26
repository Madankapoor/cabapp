package passenger

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/models"
	passengerservice "github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/service/passenger"
	"github.com/gin-gonic/gin"
)

// Passenger type alias
type Passenger = models.Passenger

func create(c *gin.Context) {
	var passenger models.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		c.AbortWithStatus(400)
		return
	}
	passengerservice.Create(c, &passenger)
	c.JSON(200, passenger)
}

func list(c *gin.Context) {
	passengers, err := passengerservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, passengers)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	passenger, err := passengerservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, passenger)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := passengerservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var passenger models.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		c.AbortWithStatus(400)
		return
	}
	passenger, err := passengerservice.Update(c, &passenger)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, passenger)
}
