package models

import (
	"github.com/jinzhu/gorm"
)

// CabDriver ...
type CabDriver struct {
	gorm.Model
	Email    string `gorm:"Column:email;UNIQUE_INDEX" json:"email" binding:"required" `
	Name     string `gorm:"Column:name" json:"name" binding:"required" `
	MobileNo string `gorm:"Column:mobileno;UNIQUE_INDEX" json:"mobileno" binding:"required"`
}
