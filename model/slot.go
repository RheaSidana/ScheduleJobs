package model

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}
