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
	data, err := Models.GetAllRecentSensorData()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetSensorDataGraph(sensorLocation string, sensorType string) (*[]Models.SensorDataGraphResponse, error) {
	data, err := Models.GetSensorDataGraph(sensorType, sensorLocation)
	dateMap := make(map[string]Models.Pair[int, int])
	response := []Models.SensorDataGraphResponse{}

	for _, value := range *data {
		date := value.CreatedAt.Format(time.DateOnly)
		pair := dateMap[date]
		pair.First += value.SensorValue
		pair.Second++
		dateMap[date] = pair
	}

	for date, v := range dateMap {

		response = append(response, Models.SensorDataGraphResponse{SensorValue: (v.First / v.Second), Date: date})
	}

	if err != nil {
		return nil, err
	}

	return &response, nil
}
