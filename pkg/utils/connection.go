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
		station := findStationByName(connection.StationOne)
		if station.Name == connection.StationOne {
			station.Connections = append(station.Connections, findStationByName(connection.StationTwo))
			models.GetStationsMap().UpdateStation(station) //! We needed to update the stationMap struct
		}
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
