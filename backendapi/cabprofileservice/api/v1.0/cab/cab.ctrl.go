package cab

import (
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	cabservice "github.com/Madankapoor/cabapp/backendapi/cabprofileservice/service/cab"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var cab models.Cab
	if err := c.BindJSON(&cab); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": false, "message": err.Error(),
		})
		return
	}
	cabservice.Create(c, &cab)
	c.JSON(200, cab)
}

func list(c *gin.Context) {
	cabs, err := cabservice.List(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cabs)
}

func read(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	cab, err := cabservice.Get(c, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, cab)
}

func remove(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	err := cabservice.Delete(c, id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(204)
}

func update(c *gin.Context) {
	var cab models.Cab
	if err := c.BindJSON(&cab); err != nil {
		c.AbortWithStatus(400)
		return
	}
	cab, err := cabservice.Update(c, &cab)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, cab)
}
