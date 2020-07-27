package models

import (
	"github.com/jinzhu/gorm"
)

// Trip ...
type Trip struct {
	gorm.Model
	PassengerID uint    `json:"passenger_id" binding:"required" `
	CabID       uint    `json:"cab_id" binding:"required" `
	PickupLat   float64 `json:"pickup_lat" binding:"required"`
	PickupLong  float64 `json:"pickup_long" binding:"required"`
	DropLat     float64 `json:"drop_lat" binding:"required"`
	DropLong    float64 `json:"drop_long" binding:"required"`
}

// TripRequest ...
type TripRequest struct {
	PassengerID uint    `json:"passenger_id" binding:"required" `
	PickupLat   float64 `json:"pickup_lat" binding:"required"`
	PickupLong  float64 `json:"pickup_long" binding:"required"`
	DropLat     float64 `json:"drop_lat" binding:"required"`
	DropLong    float64 `json:"drop_long" binding:"required"`
}

// GetTripModel ...
func (tripRequest *TripRequest) GetTripModel() *Trip {
	var trip Trip = Trip{}
	trip.PassengerID = tripRequest.PassengerID
	trip.PickupLat = tripRequest.PickupLat
	trip.PickupLong = tripRequest.PickupLong
	trip.DropLat = tripRequest.DropLat
	trip.DropLong = tripRequest.DropLong
	return &trip
}
