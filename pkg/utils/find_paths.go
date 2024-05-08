package utils

import (
	"fmt"
	"math"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func FindPath() {
	instance := models.GetInstance()
	startStation := instance.StartStation
	endStation := instance.EndStation

	// Find the first station from the start station
	firstStation := findStationByName(startStation)
	firstStation.IsStart = true
	models.StationsInstance.UpdateStation(firstStation)
	endingStation := findStationByName(endStation)
	endingStation.IsFinish = true
	models.StationsInstance.UpdateStation(endingStation)
	var path []models.Station
	var currentStation = firstStation
	path = append(path, currentStation)
	for currentStation.Name != endStation {

		// Calculate distances to connected stations
		distances := make(map[string]float64)
		for _, connectedStation := range currentStation.Connections {
			if !connectedStation.IsVisited {
				distance := FindStationConnectionsDistance(currentStation, connectedStation)
				distances[connectedStation.Name] = distance
				fmt.Println("Current station: ", currentStation.Name, "Connected station: ", connectedStation.Name, "Map: ", distances)
			}
		}

		currentStation.IsVisited = true
		models.StationsInstance.UpdateStation(currentStation)

		// Find the next closest station
		var nextClosestStationName string
		minDistance := math.Inf(1)
		for stationName, distance := range distances {

			station := findStationByName(stationName)
			if station.IsVisited {
				continue
			}

			if distance < minDistance {
				minDistance = distance
				nextClosestStationName = stationName
			}
		}

		if nextClosestStationName == "" {
			break
		}

		nextClosestStation := findStationByName(nextClosestStationName)

		// Append the next closest station to the path
		path = append(path, nextClosestStation)

		// Update current station for the next iteration
		currentStation = nextClosestStation
	}

	// Create a Path struct and add it to Paths instance
	pathStruct := models.Path{PathStations: path}
	pathsInstance := models.GetPaths()
	pathsInstance.AddPath(pathStruct)

	fmt.Println("Path:")
	for _, station := range path {
		fmt.Println(station.Name)
	}
}

func FindStationConnectionsDistance(station models.Station, connectedStation models.Station) float64 {
	var distance float64
	var distanceChange bool = false
	allConnections := models.GetConnectionsP()

	// Loop through all connections
	for _, connection := range allConnections.Connections {

		// Check if the connection matches the provided stations
		if connection.StationOne == station.Name && connection.StationTwo == connectedStation.Name {
			// If the connection matches, set the distance and indicate a change in distance
			distance = connection.Distance
			distanceChange = true
		} else if connection.StationTwo == station.Name && connection.StationOne == connectedStation.Name {
			distance = connection.Distance
			distanceChange = true
		}
		if distanceChange {
			break
		}
	}
	return distance
}

func FindTwoPaths() {
	FindPath()
	FindPath()
}
