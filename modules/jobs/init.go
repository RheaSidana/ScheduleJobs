package jobs

import "gorm.io/gorm"

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(jobsRepository Repository) Handler {
	return Handler{repository: jobsRepository}
}