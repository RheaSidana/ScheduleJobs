package main

import (
	"ScheduleJobs/initializer"
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

	r.Run()
}
