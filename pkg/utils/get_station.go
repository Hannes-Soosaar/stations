package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addStationsToMap(line string) {

	newStation := models.Station{}
	params := strings.Split(line, ",")
	stations := models.GetStationsMap()

	if len(params) < 2 {
		fmt.Println("There are not enough params")
	}

	newStation.Name = params[0]

	x, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	y, err := strconv.Atoi(params[2])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	newStation.X = float64(x)
	newStation.Y = float64(y)

	stations.StationsMap = append(stations.StationsMap, newStation)
}

func findStationByName(name string) models.Station {
	stations := models.GetStationsMap()
	var matchingStation models.Station
	for _, station := range stations.StationsMap {
		if station.Name == name {
			matchingStation = station
		}
	}
	return matchingStation
}

func getStationCord(stationName string) []float64 {
	stations := models.GetStationsMap()
	var stationCoordinates []float64
	for _, station := range stations.StationsMap {
		if station.Name == stationName {
			stationCoordinates = append(stationCoordinates, station.X)
			stationCoordinates = append(stationCoordinates, station.Y)
		}
	}
	return stationCoordinates
}
