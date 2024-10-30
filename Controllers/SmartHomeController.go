package controllers

import (
	Models "RCON_Server/Models"
	services "RCON_Server/Services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpenGarage(c *gin.Context) {
	err := services.OpenGarage()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Response": err.Error()})
		return
	}
}

func SaveReading(c *gin.Context) {
	var request = new(Models.SensorRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
		return
	}

	err := services.SaveSensorReading(*request)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "Successfully Saved."})
}

func GetSensorValues(c *gin.Context) {
	values, err := services.GetSensorValues()

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": values})
}
