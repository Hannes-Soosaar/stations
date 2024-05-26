package utils

import (
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
	start := FindStationByName(instance.StartStation)
	start.IsVisited = true
	start.IsStart = true
	models.StationsInstance.UpdateStation(start)
	end := FindStationByName(instance.EndStation)
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

			rout := models.Rout{
				StationNames: path,
			}
			models.GetRouts().AddRoutToRouts(rout)
			return path, true
		}
		for _, connectedStation := range FindStationByName(current.Name).Connections {
			neighbor := FindStationByName(connectedStation.Name)
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
				station := FindStationByName(pathNames[i])
				nextStation := FindStationByName(pathNames[i+1])
				station.RemoveConnection(nextStation.Name) // ? why are we removing the Connection ?
				models.StationsInstance.UpdateStation(station)
			}
			var path []models.Station
			for _, stationName := range pathNames {
				station := FindStationByName(stationName)
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
	instance := models.GetInstance()
	if numOfPaths == 0 {
		err := fmt.Errorf("error: no path exists between the start (%s) and end (%s) stations.", instance.StartStation, instance.EndStation)
		fmt.Println(err)
		os.Exit(1)
	}
	_, trainsPerPath := simulateTurns(allPossiblePaths.Paths)
	simulateTurnsHS2(trainsPerPath)
}

func simulateTurns(paths []models.Path) (int, []int) {
	instance := models.GetInstance()
	trainAmount := instance.NumberOfTrains
	_ = instance.EndStation
	trains := models.GetTrains()
	_ = models.GetRouts().Routs
	numOfPaths := len(paths)
	tempCount := len(trains.Trains)
	for tempCount > 0 {
		for _, train := range trains.Trains {
			if train.CurrentStation == instance.StartStation {
				models.GetTrains().UpdateTrainLocation(train.Id, GetNextStationOnPath(train.CurrentStation, 0))
			} else if train.CurrentStation == instance.EndStation {
				models.GetTrains().RemoveTrainById(train.Id)
			}
			if train.CurrentStation == instance.EndStation {
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

	for !allTrainsAtDestination {
		result = ""
		for _, train1 := range trains.Trains {
			if !train1.IsAtDestination {
				nextStation := GetNextStationOnPath(train1.CurrentStation, train1.TrainOnRout)
				if stationMap, exists := routeStationsMap[train1.TrainOnRout]; exists {
					if val, ok := stationMap[nextStation]; ok && !val {
						routeStationsMap[train1.TrainOnRout][nextStation] = true
						routeStationsMap[train1.TrainOnRout][train1.CurrentStation] = false
						routeStationsMap[train1.TrainOnRout][instance.StartStation] = false
						models.GetTrains().UpdateTrainLocation(train1.Id, nextStation)
						if train1.CurrentStation != instance.StartStation {
							result += "T" + strconv.Itoa(train1.Id+1) + "-" + train1.CurrentStation + " "
						}
						if nextStation == instance.EndStation {
							models.GetTrains().UpdateTrainLocation(train1.Id, nextStation)
							models.GetTrains().SetArrivedAtDestinationById(train1.Id)
						}
					}
				}
			} else if train1.IsAtDestination && !train1.DestinationPrinted {
				result += "T" + strconv.Itoa(train1.Id+1) + "-" + train1.CurrentStation + " "
				models.GetTrains().SetDestinationPrintedById(train1.Id)
				routeStationsMap[train1.TrainOnRout][instance.EndStation] = false
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

// ? is this function used
func FindStationConnectionsDistance(station models.Station, connectedStation models.Station) float64 {
	var distance float64
	var distanceChange bool = false
	allConnections, err := models.GetConnectionsP()
	if err != nil {
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

func GetShortestPath(trainID int) string {
	currentStation := FindStationByName(findCurrentStationName(trainID))
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

// useful for debug

// func displayPaths() {
// 	routs := models.GetRouts()
// 	for i, path := range routs.Routs {
// 		fmt.Printf("Currently at path %d \n ", i)
// 		fmt.Println(path)
// 	}
// }

// func waitForKeyPress() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print(" ...")
// 	reader.ReadString('\n')
// }
