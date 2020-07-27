package cabtype

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	cabtypeservice "github.com/Madankapoor/cabapp/backendapi/cabprofileservice/service/cabtype"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var passenger models.CabType
	if err := c.BindJSON(&passenger); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cabtypeservice.Create(c, &passenger)
	c.JSON(200, passenger)
}

func list(c *gin.Context) {
	cabs, err := cabtypeservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cabs)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	cab, err := cabtypeservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, cab)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := cabtypeservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var cabtype models.CabType
	if err := c.BindJSON(&cabtype); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cabtype, err := cabtypeservice.Update(c, &cabtype)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cabtype)
}
