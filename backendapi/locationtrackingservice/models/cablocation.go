package models

import (
	"github.com/jinzhu/gorm"
)

// CabLocation ...
type CabLocation struct {
	gorm.Model
	Lat      float64 `json:"lat" binding:"required"`
	Long     float64 `json:"long" binding:"required"`
	MortonNo float64 `json:"morton_number"`
	CabID    uint    `gorm:"UNIQUE_INDEX" json:"cab_id" binding:"required"`
}
