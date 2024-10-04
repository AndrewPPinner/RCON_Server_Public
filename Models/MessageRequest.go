package models

type MessageRequest struct {
	Message string `binding:"required"`
}
