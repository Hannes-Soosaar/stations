package utils

import (
	"fmt"
	"os"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func mapConnections(cs []string) {
	var connection models.Connection
	connections, err := models.GetConnectionsP()
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range cs {
		split := strings.Split(c, "-")
		if len(split) == 2 {
			connection.StationOne = split[0]
			connection.StationTwo = split[1]
		} else {
			fmt.Println("not a valid connection")
		}
		connections.Connections = append(connections.Connections, connection)
	}

	for i, connection1 := range connections.Connections {
		for j, connection2 := range connections.Connections {
			if i != j {
				if connection1.StationOne == connection2.StationOne && connection1.StationTwo == connection2.StationTwo {
					err := fmt.Errorf("error: Duplicate routs exists between %s and %s", connection1.StationOne, connection1.StationTwo)
					fmt.Println(err)
					os.Exit(1)
				} else if connection1.StationOne == connection2.StationTwo && connection1.StationTwo == connection2.StationOne {
					err := fmt.Errorf("error: Duplicate reversed routes exist between between %s and %s", connection1.StationOne, connection1.StationTwo)
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

func getConnections() {

	allConnections, err := models.GetConnectionsP()
	if err != nil {
		fmt.Println(err)
	}
	for _, connection := range allConnections.Connections {
		stationOne := FindStationByName(connection.StationOne)
		stationTwo := FindStationByName(connection.StationTwo)
		if stationOne.Name == connection.StationOne {
			stationOne.Connections = append(stationOne.Connections, FindStationByName(connection.StationTwo))
			stationOne.ConnObj = append(stationOne.ConnObj, connection)
			stationTwo.Connections = append(stationTwo.Connections, FindStationByName(connection.StationOne))
			stationTwo.ConnObj = append(stationTwo.ConnObj, connection)
			models.GetStationsMap().UpdateStation(stationOne)
			models.GetStationsMap().UpdateStation(stationTwo)
		} else {

		}
	}
}

func StationInConnectionIsAStation(){
	connections, _ := models.GetConnectionsP() 
for _,connection := range connections.Connections{
		if !StationExistByName(connection.StationOne){
			err := fmt.Errorf("error: a connection is made with %s station which does not exist",connection.StationOne)
			fmt.Println(err)
			os.Exit(1)
		}else if  !StationExistByName(connection.StationTwo){
			err := fmt.Errorf("error: a connection is made with %s station which does not exist",connection.StationTwo)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
// DECIDED NOT TO USE DISTANCE AS IT WAS NOT NEEDED
// func AddDistanceToConnection() {
// 	allConnections, err := models.GetConnectionsP()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	deltaCordSqr := make([]float64, 2)
// 	for i, connection := range allConnections.Connections {
// 		stationOneCord := getStationCord(connection.StationOne)
// 		stationTwoCord := getStationCord(connection.StationTwo)
// 		deltaCordSqr[0] = math.Pow(stationOneCord[0]-stationTwoCord[0], 2)
// 		deltaCordSqr[1] = math.Pow(stationOneCord[1]-stationTwoCord[1], 2)
// 		distBetweenStations := math.Sqrt(deltaCordSqr[0] + deltaCordSqr[1])
// 		allConnections.Connections[i].Distance = distBetweenStations
// 	}
// 	AddConnectionToStations()
// }
