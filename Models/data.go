package models

import (
	token "RCON_Server/Utils"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"size:255;not null;unique"`
	HashedPassword string `gorm:"size:255;not null"`
}

type SensorData struct {
	gorm.Model
	SensorValue int `gorm:"not null"`
	//Could probably normalize this but, eh
	SensorLocation string `gorm:"size:255;not null"`
	SensorType     string `gorm:"size:255;not null"`
}

func (authRequest *AuthRequest) CreateUser() error {
	u := User{Username: authRequest.Username, HashedPassword: authRequest.Password}
	err := Database.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (authRequest *AuthRequest) AttempLogin() (*LoginResponse, error) {
	user := User{}
	res := LoginResponse{}

	err := Database.Model(&User{}).Where("username = ?", authRequest.Username).Take(&user).Error

	if err != nil {
		res.Message = "Login attempt failed."
		return &res, err
	}

	err = VerifyPassword(authRequest.Password, user.HashedPassword)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		res.Message = "Login attempt failed."
		return &res, err
	}

	res.Token, err = token.GenerateToken(user.ID)

	if err != nil {
		res.Message = "Failed to generate new token. Please try again later"
		return &res, err
	}

	res.Message = "Sucessful Login. Welcome back."
	return &res, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// gorm hook function. Gets called before every User model save
func (user *User) BeforeSave(tx *gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.HashedPassword = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func SaveReading(request SensorRequest) error {
	data := SensorData{SensorValue: request.SensorValue, SensorLocation: request.Location, SensorType: request.SensorType}

	err := Database.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAllRecentSensorData() (*[]SensorData, error) {
	data := []SensorData{}
	err := Database.Raw("SELECT distinct on (sensor_type, sensor_location) * FROM sensor_data order by sensor_type, sensor_location, created_at desc;").Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GetSensorDataGraph(sensorType string, sensorLocation string) (*[]SensorData, error) {
	data := []SensorData{}
	err := Database.Raw("SELECT * FROM sensor_data where sensor_location = ? AND sensor_type = ? order by created_at desc;", sensorLocation, sensorType).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}
