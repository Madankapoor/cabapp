package passengerlocation

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/models"
	passengerlocationservice "github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/service/passengerlocation"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var passengerlocation models.PassengerLocation
	if err := c.BindJSON(&passengerlocation); err != nil {
		c.AbortWithStatus(400)
		return
	}
	passengerlocationservice.Create(c, &passengerlocation)
	c.JSON(200, passengerlocation)
}

func list(c *gin.Context) {
	passengerlocations, err := passengerlocationservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, passengerlocations)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	passengerlocation, err := passengerlocationservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, passengerlocation)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := passengerlocationservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var passengerlocation models.PassengerLocation
	if err := c.BindJSON(&passengerlocation); err != nil {
		c.AbortWithStatus(400)
		return
	}
	passengerlocation, err := passengerlocationservice.Update(c, &passengerlocation)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, passengerlocation)
}
