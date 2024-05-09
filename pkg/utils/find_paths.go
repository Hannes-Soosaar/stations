package utils

import (
	"fmt"
	"log"
	"math"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func FindPath() bool {
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

	// if the last found station of this path is not the required end station, this code will not add the path and will exit the function.
	lastStation := path[len(path)-1]
	if lastStation.Name != endStation {
		newUniquePathFound := false
		return newUniquePathFound
	}

	// Create a Path struct
	pathStruct := models.Path{PathStations: path}
	pathsInstance := models.GetPaths()
	MoveTrains() //? This is HS trial function
	// if there's at least 1 unique path found, this code will check if the next path is identical to the previous one.
	// If it is then it will not add it to the list of paths and the function will return since all unique paths have been found.
	if pathsInstance.Paths != nil {
		if len(pathStruct.PathStations) == len(pathsInstance.Paths[len(pathsInstance.Paths)-1].PathStations) {
			equal := true
			for i, station := range pathStruct.PathStations {
				if station.Name != pathsInstance.Paths[len(pathsInstance.Paths)-1].PathStations[i].Name {
					equal = false
					break
				}
			}

			if equal {
				newUniquePathFound := false
				return newUniquePathFound
			}
		}
	}

	//if all checks are passed that means an unique path has been found so it will be added here.
	pathsInstance.AddPath(pathStruct)

	fmt.Println("Path:")
	for _, station := range path {
		fmt.Println(station.Name)
	}

	newUniquePathFound := true
	return newUniquePathFound
}

// TODO only pass in strings.
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

func FindAllUniquePaths() {

	newUniquePathFound := true

	for newUniquePathFound {
		newUniquePathFound = FindPath()
	}
}

func GetShortestPath(trainID int) string {
	currentStation := findStationByName(findCurrentStationName(trainID))
	var trainToMoveTo string

	log.Println(currentStation.ConnObj) // TODO the ConnObj does not contain distance

	var distance float64
	for _, stationConnections := range currentStation.ConnObj {
		if stationConnections.Distance == 0 || stationConnections.Distance < distance {
			distance = stationConnections.Distance
			if currentStation.Name == stationConnections.StationOne {
				trainToMoveTo = stationConnections.StationTwo
			} else if currentStation.Name == stationConnections.StationTwo {
				trainToMoveTo = stationConnections.StationTwo
			} else {
				fmt.Println("Something is off!")
			}
		}
	}
	return trainToMoveTo
}

// logic is to find the next station that is the shortest distance away
func findClosestStation(connection models.Connections) {

}
