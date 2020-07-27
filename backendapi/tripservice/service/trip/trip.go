package trip

import (
	"github.com/Madankapoor/cabapp/backendapi/tripservice/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CabLocation is a simplified CabLocation
type CabLocation struct {
	ID    int     `json:"id"`
	Lat   float64 `json:"lat" `
	Long  float64 `json:"long" `
	CabID uint    `json:"cab_id"`
}

// LocationServiceError generic error message
type LocationServiceError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// BookTrip object
func BookTrip(c *gin.Context, trip *models.Trip) (*models.Trip, error) {
	db := c.MustGet("db").(*gorm.DB)
	locationservice := c.MustGet("locationservice").(*sling.Sling)
	var cablocations []CabLocation = []CabLocation{}
	var locationerror LocationServiceError = LocationServiceError{}
	_, err := locationservice.New().Get("/api/v1.0/cablocation").Receive(&cablocations, &locationerror)
	if err != nil {
		return trip, err
	}
	// Algorithm to pick cablocation needs to decided. For now just taking first element.
	if len(cablocations) > 0 {
		trip.CabID = cablocations[0].CabID
	}
	// Update cabprofile that cab is booked for now and cannot be booked again.
	db.NewRecord(trip)
	db.Create(&trip)
	return trip, nil
}

// List Trips
func List(c *gin.Context) ([]models.Trip, error) {
	db := c.MustGet("db").(*gorm.DB)
	var trips []models.Trip = []models.Trip{}
	if err := db.Order("id asc").Find(&trips).Error; err != nil {
		return trips, err
	}
	return trips, nil
}

// Get Trips for given passenger_id
func Get(c *gin.Context, id int) ([]models.Trip, error) {
	db := c.MustGet("db").(*gorm.DB)
	var trips []models.Trip
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("passenger_id = ?", id).Find(&trips).Error; err != nil {
		return trips, err
	}
	return trips, nil
}

// Update Trip
func Update(c *gin.Context, trip *models.Trip) (models.Trip, error) {
	db := c.MustGet("db").(*gorm.DB)
	db.Save(&trip)
	return *trip, nil
}
