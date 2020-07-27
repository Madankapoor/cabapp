package cablocation

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/models"
	cablocationservice "github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/service/cablocation"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var cablocation models.CabLocation
	if err := c.BindJSON(&cablocation); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cablocationservice.Create(c, &cablocation)
	c.JSON(200, cablocation)
}

func list(c *gin.Context) {
	cablocations, err := cablocationservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cablocations)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	cablocation, err := cablocationservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, cablocation)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := cablocationservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var cablocation models.CabLocation
	if err := c.BindJSON(&cablocation); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cablocation, err := cablocationservice.Update(c, &cablocation)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cablocation)
}
