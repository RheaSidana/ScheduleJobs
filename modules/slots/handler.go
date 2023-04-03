package slots

import (
	"ScheduleJobs/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
}

func (h *Handler) CreateSlotHandler(c *gin.Context) {
	var newSlot model.SlotBody
	c.BindJSON(&newSlot)

	if newSlot == (model.SlotBody{}){
		c.JSON(400, ErrorResponse{Message: "Unable to add slot. Error: "})
		return
	}

	slot, err := h.repository.Create(newSlot)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to add slot. Error: " + err.Error()})
		return
	}

	c.JSON(200, SlotResponse{
		Message: slot.Start.Format("2006-01-02 15:04:05") + 
		" to " + slot.End.Format("2006-01-02 15:04:05") + 
		" added"})
}

func (h *Handler) FindSlotHandler(c *gin.Context) {
	var getSlotRequest GetSlotRequest
	c.BindJSON(&getSlotRequest)

	slots, err := h.repository.FindAllCurrentSlots(getSlotRequest.UserID)
	if err != nil {
		c.JSON(400, ErrorResponse{Message: "Unable to find slots. Error: " + err.Error()})
		return
	}

	message := ""
	for i, slot := range slots{
		message += "Slot " + strconv.Itoa(i+1) + ": " + slot.Start.Weekday().String() + " " +
		slot.Start.Format("2006-01-02 15:04:05") + " to " + slot.End.Weekday().String() + " " +
		slot.End.Format("2006-01-02 15:04:05") 
		if i < len(slots)-1 {
			message += "       "
		}
	}

	c.JSON(200, SlotResponse{Message: message})
}