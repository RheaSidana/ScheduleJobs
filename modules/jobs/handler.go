package jobs

import (
	"ScheduleJobs/model"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
}

func (h *Handler) CreateJobHandler(c *gin.Context) {
	var newJob model.Job
	c.BindJSON(&newJob)
	if newJob == (model.Job{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add job."})
		return
	}

	job, err := h.repository.Create(newJob)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to add job. Error: "+err.Error()})
		return
	}

	c.JSON(200, JobResponse{Message: job.Name + " created successfully!!"})
}

// func (h *Handler) FindJobHandler(c *gin.Context) (model.Job,error){
// 	jobID, err := strconv.Atoi(c.Param("job_id"))
// 	if err != nil {
// 		c.JSON(400, JobResponse{Message: "Error: " + err.Error()})
// 		return model.Job{}, err
// 	}

// 	jobSlots, er := h.repository.FindJob(jobID)
// 	if er != nil {
// 		c.JSON(400, JobResponse{Message: "Error: " + er.Error()})
// 		return model.Job{}, err
// 	}

// 	return jobSlots,nil
// }
