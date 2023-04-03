package slots

import (
	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(slotRepository Repository) Handler {
	return Handler{repository: slotRepository}
}
