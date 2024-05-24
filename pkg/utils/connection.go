package utils

//TODO: Rename  file as it  is named wrong. It gets connections not connection
import (
	"fmt"
	"math"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func mapConnections(cs []string) {
	var connection models.Connection
	connections := models.GetConnectionsP()
	for _, c := range cs {
		split := strings.Split(c, "-")
		if len(split) == 2 {
			connection.StationOne = split[0]
			connection.StationTwo = split[1]
		} else {
			fmt.Println("not a valid connection")
		}
		connections.Connections = append(connections.Connections, connection) // initializes the first struct?
	}
}

func getConnections() {
	allConnections := models.GetConnectionsP()
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
		}
	}
}
// This function is not used currently for anything practical but might be useful if there is further development done with the programm.
func AddDistanceToConnection() {
	allConnections := models.GetConnectionsP()
	fmt.Printf("error: %v\n", error)
	deltaCordSqr := make([]float64, 2)
	for i, connection := range allConnections.Connections {
		stationOneCord := getStationCord(connection.StationOne)
		stationTwoCord := getStationCord(connection.StationTwo)
		deltaCordSqr[0] = math.Pow(stationOneCord[0]-stationTwoCord[0], 2)
		deltaCordSqr[1] = math.Pow(stationOneCord[1]-stationTwoCord[1], 2)
		distBetweenStations := math.Sqrt(deltaCordSqr[0] + deltaCordSqr[1])
		allConnections.Connections[i].Distance = distBetweenStations
	}
	addConnectionToStations()
}
