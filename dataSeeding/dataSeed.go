package main

import (
	"ScheduleJobs/dataSeeding/dataDump"
	"ScheduleJobs/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	dataDump.UserData()
}
