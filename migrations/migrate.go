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
}
