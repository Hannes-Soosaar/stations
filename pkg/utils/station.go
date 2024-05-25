package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addStationsToMap(line string) {
	newStation := models.Station{}
	params := strings.Split(line, ",")
	stations := models.GetStationsMap()
	if len(params) != 3 {
		err := fmt.Errorf("The line %s in stations is invalid. A station must have a valid name and valid coordinates example: \" station_name,1,1 \"", line)
		fmt.Println(err)
		os.Exit(1)
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
	checkStationCoordinates(x, y, newStation.Name)
	checkStationName(newStation.Name)
	stations.StationsMap = append(stations.StationsMap, newStation)
}

func FindStationByName(name string) models.Station {
	stations := models.GetStationsMap()
	var matchingStation models.Station
	for _, station := range stations.StationsMap {
		if station.Name == name {
			matchingStation = station
		}
	}
	return matchingStation
}
func StationExistByName(name string) bool {
	stations := models.GetStationsMap()
	for _, station := range stations.StationsMap {
		if station.Name == name {
			return true
		}
	}
	return false
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

func checkForDuplicateCoordinates() {
	stations := models.GetStationsMap()
	for i, station1 := range stations.StationsMap {
		for j, station2 := range stations.StationsMap {
			if i != j {
				if station1.Name == station2.Name {
					err := fmt.Errorf("error: there are duplicate station entries for %s ", station1.Name)
					fmt.Println(err)
					os.Exit(1)
				} else if station1.X == station2.X && station1.Y == station2.Y {
					err := fmt.Errorf("error: %s station and %s station are on the same coordinates (X:%.0f,Y:%.0f)", station1.Name, station2.Name, station1.X, station1.Y)
					fmt.Println(err)
					os.Exit(1)
				}

			}
		}
	}
}

func checkStationName(s string) {

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' || char == '_' {
			continue
		} else {
			err := fmt.Errorf("error: the station %s name is invalid. A station name can only contain lower case letters, positive integers and  an '_' ", s)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func checkStationCoordinates(x int, y int, stationName string) {
	if x < 0 || y < 0 {
		err := fmt.Errorf("error: Station: %s coordinates (X: %d ,Y: %d) are not valid. Coordinates must be positive integers or 0 ", stationName, x, y)
		fmt.Println(err)
		os.Exit(1)
	}
}
