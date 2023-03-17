package users

import (
	"ScheduleJobs/model"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	repository Repository
}

func (h *Handler) CreateUserHandler(c *gin.Context){
	var newUser model.User
	c.BindJSON(&newUser)
	if newUser == (model.User{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add user."})
		return
	}

	user, err := h.repository.Create(newUser)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to add user."})
		return
	}

	c.JSON(200, UserResponse{Message: user.Name + " created successfully!!"})
}

// func (h *Handler) FindUser(c *gin.Context) (model.User, error){
// 	var user int
// 	c.BindJSON(&user)
// 	return h.repository.Find(user)
// }