package passengerlocation

import (
	"github.com/Madankapoor/cabapp/backendapi/locationtrackingservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create passenger object
func Create(c *gin.Context, passengerLocation *models.PassengerLocation) {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(passengerLocation)
	db.Create(&passengerLocation)
}

// List Passengers
func List(c *gin.Context) ([]models.PassengerLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	var passengers []models.PassengerLocation = []models.PassengerLocation{}
	if err := db.Order("id asc").Find(&passengers).Error; err != nil {
		return passengers, err
	}
	return passengers, nil
}

// Get Passenger
func Get(c *gin.Context, id int) (models.PassengerLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	var passengerLocation models.PassengerLocation
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&passengerLocation).Error; err != nil {
		return passengerLocation, err
	}
	return passengerLocation, nil
}

// Delete Passenger
func Delete(c *gin.Context, id int) error {
	db := c.MustGet("db").(*gorm.DB)
	passengerLocation, err := Get(c, id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&passengerLocation)
	return nil
}

// Update passenger
func Update(c *gin.Context, passengerLocation *models.PassengerLocation) (models.PassengerLocation, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&passengerLocation)
	return *passengerLocation, nil
}
