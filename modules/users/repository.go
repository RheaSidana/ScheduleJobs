package users

import (
	"ScheduleJobs/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user model.User) (model.User, error)
	Find(userId int) (model.User, error)
}

type repository struct{
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(user model.User) (model.User, error) {
	result := r.client.Create(&user)
	
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repository) Find(userId int) (model.User, error) {
	var user model.User
	res := r.client.Where("id=?", userId).Find(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}
