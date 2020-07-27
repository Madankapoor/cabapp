package cablocation

import (
	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create cablocation object
func Create(c *gin.Context, cablocation *models.CabLocation) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(cablocation)
	db.Create(&cablocation)
}

// List cablocations
func List(c *gin.Context) ([]models.CabLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cablocations []models.CabLocation = []models.CabLocation{}
	if err := db.Order("id asc").Find(&cablocations).Error; err != nil {
		return cablocations, err
	}
	return cablocations, nil
}

// Get cablocation
func Get(c *gin.Context, id int) (models.CabLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cablocation models.CabLocation
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&cablocation).Error; err != nil {
		return cablocation, err
	}
	return cablocation, nil
}

// Delete cablocation
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	cablocation, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&cablocation)
	return nil
}

// Update cablocation
func Update(c *gin.Context, cablocation *models.CabLocation) (models.CabLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&cablocation)
	return *cablocation, nil
}
