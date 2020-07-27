package models

import (
	"github.com/jinzhu/gorm"
)

// CabType ...
type CabType struct {
	gorm.Model
	Name string `gorm:"Column:name:UNIQUE_INDEX" json:"name" binding:"required" `
}
