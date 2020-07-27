package cabtype

import (
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create cab object
func Create(c *gin.Context, cab *models.CabType) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(cab)
	db.Create(&cab)
}

// List cab drivers
func List(c *gin.Context) ([]models.CabType, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cabtypes []models.CabType = []models.CabType{}
	if err := db.Order("id asc").Find(&cabtypes).Error; err != nil {
		return cabtypes, err
	}
	return cabtypes, nil
}

// Get cab
func Get(c *gin.Context, id int) (models.CabType, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cabtype models.CabType
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&cabtype).Error; err != nil {
		return cabtype, err
	}
	return cabtype, nil
}

// Delete cab
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	cabtype, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&cabtype)
	return nil
}

// Update cab
func Update(c *gin.Context, cabtype *models.CabType) (models.CabType, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&cabtype)
	return *cabtype, nil
}
