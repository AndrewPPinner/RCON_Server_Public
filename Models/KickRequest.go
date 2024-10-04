package models

type KickRequest struct {
	Message string `binding:"required"`
	SteamID string `binding:"required"`
}
