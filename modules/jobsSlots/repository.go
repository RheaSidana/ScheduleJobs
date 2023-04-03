package jobsSlots

import (
	"ScheduleJobs/model"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user_job_slot model.User_job_slot) (model.User_job_slot, error)
	FindAll(user int) ([]model.User_job_slot, error)
}

type repository struct{
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(user_job_slot model.User_job_slot) (model.User_job_slot, error) {
	result := r.client.Create(&user_job_slot)
	
	if result.Error != nil {
		return model.User_job_slot{}, result.Error
	}

	return user_job_slot, nil
}

func (r *repository) FindAll(userID int) ([]model.User_job_slot, error){
	var jobSlots []model.User_job_slot
	res := r.client.Where("user_id=?", userID).Find(&jobSlots)

	if res.RowsAffected < 1{
		return []model.User_job_slot{}, errors.New("no slots found")
	}
	
	return jobSlots, nil
}