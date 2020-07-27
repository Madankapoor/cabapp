package cabdriver

import (
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create cab object
func Create(c *gin.Context, cab *models.CabDriver) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(cab)
	db.Create(&cab)
}

// List cab drivers
func List(c *gin.Context) ([]models.CabDriver, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cabdrivers []models.CabDriver = []models.CabDriver{}
	if err := db.Order("id asc").Find(&cabdrivers).Error; err != nil {
		return cabdrivers, err
	}
	return cabdrivers, nil
}

// Get cab
func Get(c *gin.Context, id int) (models.CabDriver, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cabdriver models.CabDriver
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&cabdriver).Error; err != nil {
		return cabdriver, err
	}
	return cabdriver, nil
}

// Delete cab
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	cabdriver, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&cabdriver)
	return nil
}

// Update cab
func Update(c *gin.Context, cabdriver *models.CabDriver) (models.CabDriver, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&cabdriver)
	return *cabdriver, nil
}
