package main

import (
	"ScheduleJobs/initializer"
	"ScheduleJobs/model"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.Db.AutoMigrate(&model.User{})
	initializer.Db.AutoMigrate(&model.Job{})
	initializer.Db.AutoMigrate(&model.Slot{})
	initializer.Db.AutoMigrate(&model.User_job_slot{})
}
