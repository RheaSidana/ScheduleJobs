package jobs

import (
	"ScheduleJobs/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(job model.Job) (model.Job, error)
	Find(jobID int) (model.Job, error)
}

type repository struct{
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(job model.Job) (model.Job, error) {
	result := r.client.Create(&job)

	if result.Error != nil {
		return model.Job{}, result.Error
	}
	return job, nil
}

func (r *repository) Find(jobID int) (model.Job, error){	
	var job model.Job
	if res := r.client.Where("id = ?", jobID).First(&job); res.Error != nil{
		return model.Job{}, res.Error
	}

	return job, nil
}