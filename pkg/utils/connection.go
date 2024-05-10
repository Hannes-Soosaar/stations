package utils

// the file is named wrong. It gets connections not connection
import (
	"fmt"
	"log"
	"math"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)
//? Reads in the Map
func mapConnections(cs []string) {
	var connection models.Connection
	//This creates the empty connections structure at the pointer
	connections := models.GetConnectionsP()
	for _, c := range cs {
		split := strings.Split(c, "-")
		if len(split) == 2 {
			connection.StationOne = split[0]
			connection.StationTwo = split[1]
		} else {
			fmt.Println("not a valid connection")
		}
		connections.Connections = append(connections.Connections, connection)// initializes the first struct?
	}
	
}

func getConnections(stationsMap models.StationsMap) models.StationsMap {
	allConnections := models.GetConnectionsP()
	fmt.Println(allConnections) // no distances!
	for _, connection := range allConnections.Connections {
		stationOne := findStationByName(connection.StationOne)
		stationTwo := findStationByName(connection.StationTwo)
		// fmt.Println("First station: ", stationOne.Name, "Second station: ", stationTwo.Name)
		if stationOne.Name == connection.StationOne {
			stationOne.Connections = append(stationOne.Connections, findStationByName(connection.StationTwo))
			stationOne.ConnObj = append(stationOne.ConnObj, connection) //! testing the direct Connection slice
			stationTwo.Connections = append(stationTwo.Connections, findStationByName(connection.StationOne))
			stationTwo.ConnObj = append(stationTwo.ConnObj, connection)	//! testing the direct connection slice 
			models.GetStationsMap().UpdateStation(stationOne) //! We needed to update the stationMap struct
			models.GetStationsMap().UpdateStation(stationTwo)
			// fmt.Println("Station name: ", stationOne.Name, "connection: ", connection.StationTwo)
		}
		// fmt.Println("FAR STATION>>> ", findStationByName("far").Connections)
	}
	return stationsMap
}

func AddDistanceToConnection() {
	allConnections := models.GetConnectionsP()
	log.Println(allConnections) // no distances
	deltaCordSqr := make([]float64, 2)
	for i, connection := range allConnections.Connections {
		stationOneCord := getStationCord(connection.StationOne)
		stationTwoCord := getStationCord(connection.StationTwo)
		deltaCordSqr[0] = math.Pow(stationOneCord[0]-stationTwoCord[0], 2)
		deltaCordSqr[1] = math.Pow(stationOneCord[1]-stationTwoCord[1], 2)
		distBetweenStations := math.Sqrt(deltaCordSqr[0] + deltaCordSqr[1])
		allConnections.Connections[i].Distance = distBetweenStations
		fmt.Printf("Distance to station  %f.2 \n",distBetweenStations)
	}
	addConnectionToStations()
	createTrains()
}
