package utils

import (
	"container/list"
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

type QueueNode struct {
	Station *models.Station
	Prev    *QueueNode
}

func FindPathWithBFS() ([]string, bool) {

	instance := models.GetInstance()

	start := findStationByName(instance.StartStation)
	start.IsVisited = true
	start.IsStart = true
	models.StationsInstance.UpdateStation(start)

	end := findStationByName(instance.EndStation)
	end.IsFinish = true
	models.StationsInstance.UpdateStation(end)

	queue := list.New()
	queue.PushBack(&QueueNode{Station: &start, Prev: nil})

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*QueueNode)
		current := node.Station

		if current.Name == instance.EndStation {
			var path []string
			for node != nil {
				path = append([]string{node.Station.Name}, path...)
				node = node.Prev
			}
			fmt.Println("Path:")
			for _, stationName := range path {
				fmt.Println(stationName)
			}
			fmt.Println("----------------------------")
			return path, true
		}

		for _, connectedStation := range findStationByName(current.Name).Connections {
			// fmt.Println("Current station: ", current.Name)
			// fmt.Println("Connected station: ", connectedStation.Name)
			// fmt.Println("Remaining connections: ", findStationByName(current.Name).Connections)
			// fmt.Println("Is connected station visited? ", connectedStation.IsVisited)
			neighbor := findStationByName(connectedStation.Name)
			// fmt.Println("neighbor station visited? ", neighbor.IsVisited)

			if !connectedStation.IsVisited {
				neighbor.IsVisited = false
			} else if connectedStation.IsVisited {
				neighbor.IsVisited = true
			}

			if connectedStation.Name == instance.StartStation {
				neighbor.IsVisited = true
			}

			if !neighbor.IsVisited {
				neighbor.IsVisited = true
				models.StationsInstance.UpdateStation(neighbor)
				// fmt.Println("Next station name and isvisited:", neighbor.Name, ",", neighbor.IsVisited)
				queue.PushBack(&QueueNode{Station: &neighbor, Prev: node})
				if queue.Len() > 250 {
					return nil, false
				}
			}
		}

	}
	return nil, false
}

func FindAllUniquePaths() {
	newUniquePathFound := true
	var pathNames []string
	for newUniquePathFound {
		pathNames, newUniquePathFound = FindPathWithBFS()
		if pathNames != nil {
			// Update stations based on the path
			for i := 0; i < len(pathNames)-1; i++ {
				station := findStationByName(pathNames[i])
				nextStation := findStationByName(pathNames[i+1])
				station.RemoveConnection(nextStation.Name)
				models.StationsInstance.UpdateStation(station)
			}

			// Add the path to the list of paths
			var path []models.Station
			for _, stationName := range pathNames {
				station := findStationByName(stationName)
				path = append(path, station)
			}
			pathStruct := models.Path{PathStations: path}
			pathsInstance := models.GetPaths()
			pathsInstance.AddPath(pathStruct)
		}
	}

	FindPathCombWithLeastTurns()
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
