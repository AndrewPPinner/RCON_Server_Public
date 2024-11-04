package main

import (
	Controllers "RCON_Server/Controllers"
	Middleware "RCON_Server/Middleware"
	Models "RCON_Server/Models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Start Server")

	Models.ConnectToDatabase()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//Use if I need to debug locally
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
	// 	AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	rcon := r.Group("RCON/api")
	rcon.POST("/login", Controllers.Login)
	//public.POST("/register", Controllers.Register) commenting out because I don't want register to be available

	protected := r.Group("RCON/api/admin")
	protected.Use(Middleware.JwtAuthMiddleware())

	protected.POST("/save", Controllers.Save)
	protected.GET("/players", Controllers.PlayerList)
	protected.POST("/ban", Controllers.BanPlayer)
	protected.POST("/broadcast", Controllers.BroadcastMessage)
	protected.POST("/kick", Controllers.KickPlayer)
	protected.GET("/status", Controllers.ServerStatus)

	smartHome := r.Group("smart/api")
	smartHome.Use(Middleware.JwtAuthMiddleware())

	smartHome.POST("/open_garage", Controllers.OpenGarage)
	smartHome.POST("/saveSensor", Controllers.SaveReading)
	smartHome.GET("/getSensorValues", Controllers.GetSensorValues)
	smartHome.POST("/getSensorGraphData", Controllers.GetSensorDataGraph)

	r.Run(":6969")

}
