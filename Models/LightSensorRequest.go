package models

type SensorRequest struct {
	Location    string  `binding:"required"`
	SensorValue float32 `binding:"required"`
	//Could make an enum but lazy
	SensorType string `binding:"required"`
}
