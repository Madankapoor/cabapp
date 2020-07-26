package passenger

import (
	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PageConstant the no of elements per page.
const PageConstant = 10

// Create passenger object
func Create(c *gin.Context, passenger *models.Passenger) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(passenger)
	db.Create(&passenger)
}

// List Passengers
func List(c *gin.Context) ([]models.Passenger, error) {
	db := c.MustGet("db").(*gorm.DB)
	var passengers []models.Passenger = []models.Passenger{}
	if err := db.Order("id asc").Find(&passengers).Error; err != nil {
		return passengers, err
	}
	return passengers, nil
}

// Get Passenger
func Get(c *gin.Context, id int) (models.Passenger, error) {
	db := c.MustGet("db").(*gorm.DB)
	var passenger models.Passenger
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&passenger).Error; err != nil {
		return passenger, err
	}
	return passenger, nil
}

// Delete Passenger
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	passenger, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&passenger)
	return nil
}

// Update passenger
func Update(c *gin.Context, passenger *models.Passenger) (models.Passenger, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&passenger)
	return *passenger, nil
}
