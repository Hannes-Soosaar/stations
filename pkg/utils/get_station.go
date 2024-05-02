package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// this is a set/map function
func getStation(line string) models.Station {

	newStation := models.Station{}
	params := strings.Split(line, ",")
	stations := models.GetStationsMap()

	if len(params) < 2 {
		fmt.Println("There are not enough params")
		return newStation
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

	newStation.X = x
	newStation.Y = y

	stations.StationsMap = append(stations.StationsMap, newStation)

	fmt.Println(stations.StationsMap)
	return newStation //redundant as its already stored in instance
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
