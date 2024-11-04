package models

type SensorGraphRequest struct {
	Location string `binding:"required"`
	Type     string `binding:"required"`
}
