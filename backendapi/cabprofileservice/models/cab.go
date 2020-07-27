package models

import (
	"github.com/jinzhu/gorm"
)

// Cab ...
type Cab struct {
	gorm.Model
	CabNo       string    `gorm:"Column:cabno;UNIQUE_INDEX" json:"cab_no" binding:"required" `
	CabTypeID   uint      `json:"cab_type_id"   binding:"required" `
	CabType     CabType   `gorm:"save_associations:false"`
	CabDriverID uint      `json:"cab_driver_id" binding:"required" `
	CabDriver   CabDriver `gorm:"save_associations:false"`
}

// CabInput ...
type CabInput struct {
	CabNo       string `json:"cab_no" binding:"required" `
	CabTypeID   uint   `json:"cab_type_id"   binding:"required" `
	CabDriverID uint   `json:"cab_driver_id" binding:"required" `
}

// GetCabModel ..
func (cabInput *CabInput) GetCabModel() *Cab {
	var cab Cab = Cab{}
	cab.CabNo = cabInput.CabNo
	cab.CabTypeID = cabInput.CabTypeID
	cab.CabDriverID = cabInput.CabDriverID
	return &cab
}
