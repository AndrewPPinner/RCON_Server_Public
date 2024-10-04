package controllers

import (
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
