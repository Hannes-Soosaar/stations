package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func FindPath() bool {
	instance := models.GetInstance()
	startStation := instance.StartStation
	endStation := instance.EndStation
	// Find the first station from the start station
	firstStation := findStationByName(startStation)
	firstStation.IsStart = true
	firstStation.IsVisited = true
	models.StationsInstance.UpdateStation(firstStation)
	endingStation := findStationByName(endStation)
	endingStation.IsFinish = true

	models.StationsInstance.UpdateStation(endingStation)

	var path []models.Station
	currentStation := firstStation

	for {
		path = append(path, currentStation)
		if currentStation.Name == endStation {
			// Path found
			pathStruct := models.Path{PathStations: path}
			pathsInstance := models.GetPaths()
			if !isUniquePath(pathsInstance, pathStruct) {
				return false
			}
			pathsInstance.AddPath(pathStruct)
			fmt.Println("----------------------------")
			fmt.Println("Path:")
			for _, station := range path {
				fmt.Println(station.Name)
			}
			firstStation.IsVisited = false
			models.StationsInstance.UpdateStation(firstStation)
			endingStation.IsVisited = false
			models.StationsInstance.UpdateStation(endingStation)
			return true
		}

		var nextStation models.Station
		found := false
		for _, connectedStation := range currentStation.Connections {
			if !findStationByName(connectedStation.Name).IsVisited {
				fmt.Println()
				nextStation = findStationByName(connectedStation.Name)
				found = true
				break
			}
		}

		if !found {
			return false // No unvisited connected stations, exit loop
		}

		nextStation.IsVisited = true
		models.StationsInstance.UpdateStation(nextStation)
		currentStation = nextStation
	}

}

func isUniquePath(pathsInstance *models.Paths, pathStruct models.Path) bool {
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
				return false
			}
		}
	}
	return true
}

func FindAllUniquePaths() {
	newUniquePathFound := true
	for newUniquePathFound {
		newUniquePathFound = FindPath()
	}
	fmt.Println("----------------------------")

	FindPathCombWithLeastTurns()
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

func GetShortestPath(trainID int) string {
	currentStation := findStationByName(findCurrentStationName(trainID))
	var trainToMoveTo string
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

func FindPathCombWithLeastTurns() {
	allPossiblePaths := models.GetPaths()
	numOfPaths := len(allPossiblePaths.Paths)
	var pathsToSimulate []models.Path
	var simulationResults []int
	count := numOfPaths

	if numOfPaths == 0 {
		fmt.Println("Error: there are no valid paths!")
	}

	for i := 0; i < numOfPaths; i++ {
		pathsToSimulate = append(pathsToSimulate, allPossiblePaths.Paths[i])
		simulationResults = append(simulationResults, simulateTurns(pathsToSimulate))
	}

	fmt.Println("Simulation results: ")
	comment := "Path %d results: "
	for j := 0; j < count; j++ {
		fmt.Printf(comment, j+1)
		fmt.Println(simulationResults[j])
	}
	// fmt.Println("First path only turns: ", simulationResults[0])
	// fmt.Println("First and second path turns: ", simulationResults[1])
	// fmt.Println("First, second and third path turns: ", simulationResults[2])
}

func simulateTurns(paths []models.Path) int {
	instance := models.GetInstance()
	trainAmount := instance.NumberOfTrains
	numOfPaths := len(paths)
	var minTurnCounts []int
	var turnCount int = 0
	//minimum turn count is the amount of turns that it takes for the first train to reach the end

	for i := 0; i < numOfPaths; i++ {
		minTurnCount := len(paths[i].PathStations) - 1
		minTurnCounts = append(minTurnCounts, minTurnCount)
		minTurnCount = 0
	}

	for trainAmount > 0 {
		turnCount++

		for i := 0; i < numOfPaths; i++ {
			if turnCount >= minTurnCounts[i] {
				trainAmount--
			}
		}
	}
	return turnCount
}
