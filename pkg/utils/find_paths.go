package utils

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"

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
	queue.PushBack(&QueueNode{Station: &start, Prev: nil}) //? what is this
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

			rout := models.Rout{
				StationNames: path,
			}
			models.GetRouts().AddRoutToRouts(rout)
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
		// fmt.Println("Paths")
		// fmt.Println(pathNames)
		if pathNames != nil {
			// Update stations based on the path
			for i := 0; i < len(pathNames)-1; i++ {
				station := findStationByName(pathNames[i])
				nextStation := findStationByName(pathNames[i+1])
				station.RemoveConnection(nextStation.Name) // ? why are we removing the Connection ?
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
	// displayPaths()
}

func FindPathCombWithLeastTurns() {
	allPossiblePaths := models.GetPaths()
	numOfPaths := len(allPossiblePaths.Paths)
	var pathsToSimulate []models.Path
	var simulationResults []int
	count := numOfPaths
	simulateTurns(allPossiblePaths.Paths)
	simulateTurnsHS2()

	if numOfPaths == 0 {
		fmt.Println("Error: there are no valid paths!")
	}
	for i := 0; i < numOfPaths; i++ {
		pathsToSimulate = append(pathsToSimulate, allPossiblePaths.Paths[i])
		simulationResults = append(simulationResults, simulateTurns(pathsToSimulate))
	}
	// fmt.Println("Simulation results: ")
	// comment := "Using the first %d path(s): %d turns will be made\n"
	for j := 0; j < count; j++ {
		// fmt.Printf(comment, j+1, simulationResults[j])
	}
}

func simulateTurns(paths []models.Path) int {
	instance := models.GetInstance()
	trainAmount := instance.NumberOfTrains
	_ = instance.EndStation
	trains := models.GetTrains()
	_ = models.GetRouts().Routs
	numOfPaths := len(paths)
	tempCount := len(trains.Trains)
	// find The next station
	// Start moving trains from the starting station to the end station
	// when a train is at the end station remove from the Trains obj.
	// continue while the trains object has trains. finish when all trains have been moved

	for tempCount > 0 { // checks to see how many trains are waiting
		// log.Println(tempCount)
		for _, train := range trains.Trains { // go through the trains
			if train.CurrentStation == instance.StartStation {
				models.GetTrains().UpdateTrainLocation(train.Id, GetNextStationOnPath(train.CurrentStation, 0))
			} else if train.CurrentStation == instance.EndStation {
				models.GetTrains().RemoveTrainById(train.Id)
			}
			if train.CurrentStation == instance.EndStation {
				// fmt.Printf("Train At finish %s", train.CurrentStation)
				models.GetTrains().RemoveTrainById(train.Id)
			}
		}
		tempCount--
	}

	var minTurnCounts []int
	var turnCount int = 0

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

func simulateTurnsHS2() {
	instance := models.GetInstance()
	trains := models.GetTrains()
	routs := models.GetRouts()
	allTrainsAtDestination := false
	// nextStation := ""
	result :=""
	turns := 1
	routeStationsMap := make(map[int]map[string]bool)

// Creates trains	
	for i := 0; i < instance.NumberOfTrains; i++ {
		train := models.Train{
			Id:             i,
			CurrentStation: instance.StartStation,
		}
		trains.Trains = append(trains.Trains, train)
	}

// Holds the maps with boolean values
	for i, rout := range routs.Routs {
		stationMap := make(map[string]bool)
		for _, station := range rout.StationNames {
			stationMap[station] = false
		}
		routeStationsMap[i] = stationMap
	}

	// temp division for routs
	for i, train := range trains.Trains {
		j := i % len(routs.Routs)
		models.GetTrains().UpdateTrainOnRout(train.Id, j)
	}

	for !allTrainsAtDestination { // checks to see how many trains are waiting
		// fmt.Println(turns)
		result =""
		for _, train1 := range trains.Trains { // go through the
			if !train1.IsAtDestination { // check if its at the end
				nextStation := GetNextStationOnPath(train1.CurrentStation, train1.TrainOnRout)
				if stationMap, exists := routeStationsMap[train1.TrainOnRout]; exists {
					if val, ok := stationMap[nextStation]; ok && !val { 
						routeStationsMap[train1.TrainOnRout][nextStation] = true            //sets the next station to occupied
						routeStationsMap[train1.TrainOnRout][train1.CurrentStation] = false // sets the current station as free
						routeStationsMap[train1.TrainOnRout][instance.EndStation] = false   // if the last station
						routeStationsMap[train1.TrainOnRout][instance.StartStation] = false // if the last station
						models.GetTrains().UpdateTrainLocation(train1.Id, nextStation)
						if train1.CurrentStation != instance.StartStation{
						result += "T "+strconv.Itoa(train1.Id)+" "+train1.CurrentStation+" "
						}
						if nextStation == instance.EndStation{
							models.GetTrains().SetArrivedAtDestinationById(train1.Id)
						}
					} else{
						// fmt.Printf("Train %d Waiting at %s, next stations %s is occupied \n", train1.Id,train1.CurrentStation,nextStation)
					}
				}
			
				// instead of 0 we will have the TrainOnRout
				// if train1.CurrentStation == instance.EndStation {
				// 	// models.GetTrains().RemoveTrainById(train1.Id)
				// 	models.GetTrains().SetArrivedAtDestinationById(train1.Id)
				// }
				// fmt.Printf("T%d-%s #On rout-%d ", train1.Id, train1.CurrentStation, train1.TrainOnRout)
			}
		}
		fmt.Println(result)
		turns++
		
		allTrainsAtDestination = checkTrainStatus()
	}
}

func checkTrainStatus() bool {
	trains := models.GetTrains()
	for _, train := range trains.Trains {
		if !train.IsAtDestination {
			return false
		}
	}
	return true
}

// Here we need to pass in
func GetNextStationOnPath(currentStationName string, routNum int) string {
	routs := models.GetRouts()
	rout := routs.Routs[routNum]
	var nextStationName string
	for i, stationName := range rout.StationNames {
		if stationName == currentStationName {
			if i < len(rout.StationNames)-1 {
				nextStationName = rout.StationNames[i+1]
			} else {
				nextStationName = rout.StationNames[len(rout.StationNames)-1]
			}
		}
	}
	return nextStationName
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

// ? Do we need this
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


func displayPaths() {
	routs := models.GetRouts()
	for i, path := range routs.Routs {
		fmt.Printf("Currently at path %d \n ", i)
		fmt.Println(path)
	}
}

func waitForKeyPress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(" ...")
	reader.ReadString('\n')
}
