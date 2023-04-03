package model

import "gorm.io/gorm"

type User_job_slot struct {
	gorm.Model
	User_Id int `json:"user_id" binding:"required"`
	Job_Id  int `json:"job_id" binding:"required"`
	Slot_Id int `json:"slot_id" binding:"required"`
}
