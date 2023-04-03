package main

import (
	"ScheduleJobs/initializer"
	"ScheduleJobs/modules/jobs"
	"ScheduleJobs/modules/jobsSlots"
	"ScheduleJobs/modules/slots"
	"ScheduleJobs/modules/users"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	users.Apis(r)
	jobs.Apis(r)
	slots.Apis(r)
	jobsSlots.Apis(r)

	r.Run()
}
