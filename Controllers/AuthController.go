package controllers

import (
	Models "RCON_Server/Models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var request = new(Models.AuthRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
		return
	}

	err := request.CreateUser()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Response": "Unable to Register User."})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "Successfully Registered."})
}

func Login(c *gin.Context) {
	var request = new(Models.AuthRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": "Bad Request"})
		return
	}

	res, err := request.AttempLogin()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Response": "Username does not exist or password is invalid."})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
