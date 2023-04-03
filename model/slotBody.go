package model

import "time"

type SlotBody struct {
	UserId     int    `json:"user_id" binding:"required"`
	JobId      int    `json:"job_id" binding:"required"`
	Start_Time string `json:"start_time" binding:"required"`
}

type SlotsToBook struct {
	User       User        `json:"user_id" binding:"required"`
	Start_Time []time.Time `json:"start_time" binding:"required"`
	Job        Job         `json:"job_id" binding:"required"`
}
