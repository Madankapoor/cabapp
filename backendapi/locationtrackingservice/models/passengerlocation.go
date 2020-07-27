package models

import (
	"github.com/jinzhu/gorm"
)

// PassengerLocation ...
type PassengerLocation struct {
	gorm.Model
	Lat         float64 `json:"lat" binding:"required"`
	Long        float64 `json:"long" binding:"required"`
	PassengerID uint    `gorm:"UNIQUE_INDEX" json:"passenger_id" binding:"required"`
}
