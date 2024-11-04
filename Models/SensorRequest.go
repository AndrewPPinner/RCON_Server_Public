package models

type SensorRequest struct {
	Location    string `binding:"required"`
	SensorValue int    `binding:"required"`
	//Could make an enum but lazy
	SensorType string `binding:"required"`
}
