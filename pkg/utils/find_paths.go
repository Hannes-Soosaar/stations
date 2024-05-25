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
			neighbor := findStationByName(connectedStation.Name)
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
			for i := 0; i < len(pathNames)-1; i++ {
				station := findStationByName(pathNames[i])
				nextStation := findStationByName(pathNames[i+1])
				station.RemoveConnection(nextStation.Name) // ? why are we removing the Connection ?
				models.StationsInstance.UpdateStation(station)
			}
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
	count := numOfPaths
	_, trainsPerPath := simulateTurns(allPossiblePaths.Paths)
	simulateTurnsHS2(trainsPerPath)

	if numOfPaths == 0 {
		fmt.Println("Error: there are no valid paths!")
	}
	for j := 0; j < count; j++ {
	}
}
// TODO: remove from 
func simulateTurns(paths []models.Path) (int, []int) {
	instance := models.GetInstance()
	trainAmount := instance.NumberOfTrains
	_ = instance.EndStation
	trains := models.GetTrains()
	_ = models.GetRouts().Routs
	numOfPaths := len(paths)
	tempCount := len(trains.Trains)

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
	var trainsOnPath = make([]int, numOfPaths)
	var trainsOnCurrentPath int

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
				trainsOnCurrentPath = trainsOnPath[i] + 1
				trainsOnPath[i] = trainsOnCurrentPath
			}

		}
	}

	for _, trainOnPath := range trainsOnPath {
		fmt.Println(trainOnPath)
	}
	fmt.Println("----------------------------")
	return turnCount, trainsOnPath
}

func simulateTurnsHS2(trainsPerPath []int) {
	instance := models.GetInstance()
	trains := models.GetTrains()
	routs := models.GetRouts()
	allTrainsAtDestination := false
	result := ""
	turns := 1
	routeStationsMap := make(map[int]map[string]bool)

	for i := 0; i < instance.NumberOfTrains; i++ {
		train := models.Train{
			Id:             i,
			CurrentStation: instance.StartStation,
		}
		trains.Trains = append(trains.Trains, train)
	}

	for i, rout := range routs.Routs {
		stationMap := make(map[string]bool)
		for _, station := range rout.StationNames {
			stationMap[station] = false
		}
		routeStationsMap[i] = stationMap
	}

	designateRoutsToTrains(trainsPerPath)

	for !allTrainsAtDestination { // checks to see how many trains are waiting
		result = ""
		for _, train1 := range trains.Trains { // go through the
			if !train1.IsAtDestination { // check if its at the end
				nextStation := GetNextStationOnPath(train1.CurrentStation, train1.TrainOnRout)
				if stationMap, exists := routeStationsMap[train1.TrainOnRout]; exists {
					if val, ok := stationMap[nextStation]; ok && !val {
						routeStationsMap[train1.TrainOnRout][nextStation] = true            //sets the next station to occupied
						routeStationsMap[train1.TrainOnRout][train1.CurrentStation] = false // sets the current station as free
						// routeStationsMap[train1.TrainOnRout][instance.EndStation] = false   //! if the last station
						routeStationsMap[train1.TrainOnRout][instance.StartStation] = false // if the last station
						models.GetTrains().UpdateTrainLocation(train1.Id, nextStation)
						if train1.CurrentStation != instance.StartStation {
							result += "T" + strconv.Itoa(train1.Id+1) + "-" + train1.CurrentStation + " " // plus one to get the trains to start form 1
						}
						if nextStation == instance.EndStation {
							models.GetTrains().UpdateTrainLocation(train1.Id, nextStation)
							models.GetTrains().SetArrivedAtDestinationById(train1.Id)
						}
					}
				}
			} else if train1.IsAtDestination && !train1.DestinationPrinted {
				result += "T" + strconv.Itoa(train1.Id+1) + "-" + train1.CurrentStation + " " // plus one to get the trains to start form 1
				models.GetTrains().SetDestinationPrintedById(train1.Id)
				routeStationsMap[train1.TrainOnRout][instance.EndStation] = false   //! if the last station
			}
		}
		fmt.Println(result)
		turns++
		allTrainsAtDestination = checkTrainStatus()
	}
}

func designateRoutsToTrains(trainsPerPath []int) {
	var trainIndex int = 0
	for !allZero(trainsPerPath) {
		for i, trainAmountForPath := range trainsPerPath {
			if trainAmountForPath <= 0 {
				continue
			} else {
				trainsPerPath[i]--
				models.GetTrains().UpdateTrainOnRout(trainIndex, i)
				trainIndex++
			}
		}
	}
}

func allZero(indexes []int) bool {
	for _, val := range indexes {
		if val != 0 {
			return false
		}
	}
	return true
}

func checkTrainStatus() bool {
	trains := models.GetTrains()
	for _, train := range trains.Trains {
		if !train.DestinationPrinted {
			return false
		}
	}
	return true
}

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
	allConnections,err := models.GetConnectionsP()
	if err !=nil {
		fmt.Println(err)
	}
	for _, connection := range allConnections.Connections {
		if connection.StationOne == station.Name && connection.StationTwo == connectedStation.Name {
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
