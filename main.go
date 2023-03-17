package main

import (
	"ScheduleJobs/initializer"
	"ScheduleJobs/modules/jobs"
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

	r.Run()
}
