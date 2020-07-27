package cab


import (
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create cab object
func Create(c *gin.Context, cab *models.Cab) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(cab)
	db.Create(&cab)
}

// List cabs
func List(c *gin.Context) ([]models.Cab, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cabs []models.Cab = []models.Cab{}
	if err := db.Preload("CabDriver").Preload("CabType").Order("id asc").Find(&cabs).Error; err != nil {
		return cabs, err
	}
	return cabs, nil
}

// Get cab
func Get(c *gin.Context, id int) (models.Cab, error) {
	db := c.MustGet("db").(*gorm.DB)
	var cab models.Cab
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&cab).Error; err != nil {
		return cab, err
	}
	return cab, nil
}

// Delete cab
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	cab, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&cab)
	return nil
}

// Update cab
func Update(c *gin.Context, cab *models.Cab) (models.Cab, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&cab)
	return *cab, nil
}
