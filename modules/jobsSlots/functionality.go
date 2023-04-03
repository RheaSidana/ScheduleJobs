package jobsSlots

import "ScheduleJobs/model"

func GetUserJobSlot(userId int, repository Repository) ([]model.User_job_slot, error) {
	jobSlotHandler := InitHandler(repository)
	return jobSlotHandler.repository.FindAll(userId)
}

func AddJobSlot(jobSlot model.User_job_slot, repository Repository) (model.User_job_slot, error) {
	jobSlotHandler := InitHandler(repository)
	return jobSlotHandler.repository.Create(jobSlot)
}
