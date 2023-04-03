package slots

import (
	"ScheduleJobs/initializer"
	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	slotHandler := InitHandler(repository)

	r.POST("/slot", slotHandler.CreateSlotHandler)
	r.GET("/slot", slotHandler.FindSlotHandler)
}
