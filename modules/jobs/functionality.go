package jobs

import "ScheduleJobs/model"

func GetJob(jobID int, repository Repository) (model.Job, error){
	jobHandler := InitHandler(repository)
	return jobHandler.repository.Find(jobID)
}