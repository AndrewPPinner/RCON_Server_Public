package controllers

import (
	Models "RCON_Server/Models"
	Services "RCON_Server/Services"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func ServerStatus(c *gin.Context) {
	options := rcon.SetDialTimeout(time.Second)
	conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"), options)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Response": "DOWN"})
		return
	}
	conn.Close()

	c.JSON(http.StatusOK, gin.H{"Response": "UP"})
}

func Save(c *gin.Context) {
	conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer conn.Close()

	_, err = conn.Execute("Save")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Response": err.Error()})
		return
	}
	//Need to find a better way. This takes 2-3 minutes to run.
	//Need to figure out why restarting a service from code SSH causes ports to not re-open
	go Services.SSHReboot()

	c.JSON(http.StatusOK, gin.H{"Response": "Saved and restart has been triggered."})
}

func PlayerList(c *gin.Context) {
	conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Response": "Unable to connect to RCON Server."})
	}
	defer conn.Close()

	response, err := conn.Execute("ShowPlayers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Response": err.Error()})
	}

	split := strings.Split(response, "\n")
	if strings.TrimSpace(split[1]) == "" {
		//Empty string because no players were found and JS is dumb
		c.JSON(http.StatusOK, gin.H{"Response": ""})
		return
	}

	var sliceRes []Models.Player
	body := strings.Split(split[1], ",")

	for i := 0; i < len(body); i += 3 {
		sliceRes = append(sliceRes, Models.Player{Name: body[i], UID: body[i+1], SteamID: body[i+2]})
	}

	c.JSON(http.StatusOK, gin.H{"Response": sliceRes})
}

func BanPlayer(c *gin.Context) {
	conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer conn.Close()
	var response = ""

	playerID, found := c.Params.Get("Steam_ID")
	if found {
		cmd := fmt.Sprintf("BanPlayer %s", playerID)
		response, err = conn.Execute(cmd)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Response": response})
}

func BroadcastMessage(c *gin.Context) {
	var request = new(Models.MessageRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
		return
	}

	go func() {
		conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		defer conn.Close()
		command := fmt.Sprintf("Broadcast %s", strings.ReplaceAll(request.Message, " ", "\xa0"))
		conn.Execute(command)
	}()

	c.JSON(http.StatusOK, gin.H{"Response": fmt.Sprintf("Broadcasted - %s", request.Message)})
}

func KickPlayer(c *gin.Context) {
	var request = new(Models.KickRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Response": err.Error()})
		return
	}

	//Display message on who was kicked
	conn, err := rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer conn.Close()

	cmd := fmt.Sprintf("Broadcast %s", request.Message)

	_, err = conn.Execute(cmd)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	time.Sleep(3 * time.Second)

	//Kick player
	conn, err = rcon.Dial("192.168.50.15:25575", os.Getenv("RCON_Password"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer conn.Close()

	cmd = fmt.Sprintf("KickPlayer %s", request.SteamID)
	response, err := conn.Execute(cmd)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": response})
}
