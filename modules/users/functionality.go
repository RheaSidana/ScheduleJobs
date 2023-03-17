package users

import (
	"ScheduleJobs/model"
)

func GetUser(id int, repository Repository) (model.User, error){
	userHandler := InitHandler(repository)
	return userHandler.repository.Find(id)
}