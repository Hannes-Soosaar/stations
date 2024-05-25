package utils

//TODO: Rename  file as it  is named wrong. It gets connections not connection
import (
	"fmt"
	// "math"
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
					err:=fmt.Errorf("error: Duplicate routs exists between %s and %s", connection1.StationOne, connection1.StationTwo)
					fmt.Println(err)
					os.Exit(1)
				} else if connection1.StationOne == connection2.StationTwo && connection1.StationTwo == connection2.StationOne {
					err:=fmt.Errorf("error: Duplicate reversed routes exist between between %s and %s", connection1.StationOne, connection1.StationTwo)
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

func getConnections() {

	allConnections, err := models.GetConnectionsP()
	stations :=models.GetStationsMap()
	var connectionHasStation bool

	for _, connection :=range allConnections.Connections{
		connectionHasStation = false
		for _, station :=range stations.StationsMap{
			if connection.StationOne == station.Name || connection.StationTwo == station.Name{
				connectionHasStation = true 
		}

		if !connectionHasStation{
					err := fmt.Errorf("error: in the %v connection connects to a station that does not exist ", connection)
					fmt.Println(err)
					os.Exit(1)
		}

	}

	if err != nil {
		fmt.Println(err)
	}
	for _, connection := range allConnections.Connections {
		stationOne := findStationByName(connection.StationOne)
		stationTwo := findStationByName(connection.StationTwo)
		if stationOne.Name == connection.StationOne {
			stationOne.Connections = append(stationOne.Connections, findStationByName(connection.StationTwo))
			stationOne.ConnObj = append(stationOne.ConnObj, connection)
			stationTwo.Connections = append(stationTwo.Connections, findStationByName(connection.StationOne))
			stationTwo.ConnObj = append(stationTwo.ConnObj, connection)
			models.GetStationsMap().UpdateStation(stationOne)
			models.GetStationsMap().UpdateStation(stationTwo)
		}else{
			
		}
	}
}
}
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
