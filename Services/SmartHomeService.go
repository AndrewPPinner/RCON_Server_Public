package services

import (
	Models "RCON_Server/Models"
	"fmt"
	"net/http"
	"time"
)

func OpenGarage() error {
	url := "http://192.168.50.46"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer res.Body.Close()
	return nil
}

func SaveSensorReading(req Models.SensorRequest) error {
	err := Models.SaveReading(req)

	if err != nil {
		return err
	}

	return nil
}

func GetSensorValues() (*[]Models.SensorData, error) {
	data, err := Models.GetSensorData()

	if err != nil {
		return nil, err
	}

	return data, nil
}
