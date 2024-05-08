package utils

// the file is named wrong. It gets connections not connection
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
		//TODO check for white spaces ?
		split := strings.Split(c, "-")
		if len(split) == 2 {
			connection.StationOne = split[0]
			connection.StationTwo = split[1]
		} else {
			fmt.Println("not a valid connection")
		}
		connections.Connections = append(connections.Connections, connection)
	}

}

func getConnections(stationsMap models.StationsMap) models.StationsMap {
	allConnections := models.GetConnectionsP()
	for _, connection := range allConnections.Connections {
		stationOne := findStationByName(connection.StationOne)
		stationTwo := findStationByName(connection.StationTwo)
		fmt.Println("First station: ", stationOne.Name, "Second station: ", stationTwo.Name)
		if stationOne.Name == connection.StationOne {
			stationOne.Connections = append(stationOne.Connections, findStationByName(connection.StationTwo))
			stationTwo.Connections = append(stationTwo.Connections, findStationByName(connection.StationOne))
			models.GetStationsMap().UpdateStation(stationOne) //! We needed to update the stationMap struct
			models.GetStationsMap().UpdateStation(stationTwo)
			fmt.Println("Station name: ", stationOne.Name, "connection: ", connection.StationTwo)
		}
		fmt.Println("FAR STATION>>> ", findStationByName("far").Connections)
	}
	return stationsMap
}

func AddDistanceToConnection() {
	allConnections := models.GetConnectionsP()
	deltaCordSqr := make([]float64, 2)
	for i, connection := range allConnections.Connections {

		stationOneCord := getStationCord(connection.StationOne)
		stationTwoCord := getStationCord(connection.StationTwo)
		deltaCordSqr[0] = math.Pow(stationOneCord[0]-stationTwoCord[0], 2)
		deltaCordSqr[1] = math.Pow(stationOneCord[1]-stationTwoCord[1], 2)
		distBetweenStations := math.Sqrt(deltaCordSqr[0] + deltaCordSqr[1])
		allConnections.Connections[i].Distance = distBetweenStations
	}
}
