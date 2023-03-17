package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Name     string `gorm:"unique" json:"name" binding:"required"` 
	Duration string `json:"duration" binding:"required"`
}
 