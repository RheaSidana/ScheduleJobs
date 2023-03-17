package jobs

import (
	"ScheduleJobs/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	jobHandler := InitHandler(repository)

	r.POST("/job", jobHandler.CreateJobHandler)
}
