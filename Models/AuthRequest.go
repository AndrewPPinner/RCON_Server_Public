package models

//Used for Login and registering
type AuthRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
