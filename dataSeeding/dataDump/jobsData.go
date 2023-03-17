package dataDump

import (
	"ScheduleJobs/initializer"
	"ScheduleJobs/model"
)

func jobsData() []model.Job{
	job1 := model.Job{
		Name: "abc",
		Duration: "3h0m0s",
	}
	job2 := model.Job{
		Name: "abcd",
		Duration: "3h0m0s",
	}
	job3 := model.Job{
		Name: "xyz",
		Duration: "3h0m0s",
	}
	job4 := model.Job{
		Name: "pqr",
		Duration: "3h0m0s",
	}
	job5 := model.Job{
		Name: "pqrt",
		Duration: "3h0m0s",
	}

	jobData := []model.Job{job1, job2, job3, job4, job5}

	return jobData
}

func addToJobsTable(jobsData []model.Job){
	for _, job := range jobsData{
		if initializer.Db.Where("name=?",job.Name).Find(&job).RowsAffected == 1 {continue}
		initializer.Db.Create(&job)
	}
}

func JobsData() {
	addToJobsTable(jobsData())
}