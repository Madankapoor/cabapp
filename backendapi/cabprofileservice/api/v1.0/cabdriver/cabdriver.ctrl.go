package cabdriver

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	cabdriverservice "github.com/Madankapoor/cabapp/backendapi/cabprofileservice/service/cabdriver"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var passenger models.CabDriver
	if err := c.BindJSON(&passenger); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cabdriverservice.Create(c, &passenger)
	c.JSON(200, passenger)
}

func list(c *gin.Context) {
	cabs, err := cabdriverservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cabs)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	cab, err := cabdriverservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, cab)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := cabdriverservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var cab models.CabDriver
	if err := c.BindJSON(&cab); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cab, err := cabdriverservice.Update(c, &cab)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cab)
}
