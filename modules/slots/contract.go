package slots

type ErrorResponse struct {
	Message string
}

type SlotResponse struct {
	Message string
}

type GetSlotRequest struct {
	UserID int `json:"user_id" binding:"required"`
}
