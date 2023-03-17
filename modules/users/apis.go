package users

import (
	"ScheduleJobs/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	userHandler := InitHandler(repository)

	r.POST("/user", userHandler.CreateUserHandler)
}
