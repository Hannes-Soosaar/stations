package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addConnectionToStations() {
	connections := models.GetConnectionsP()
	fmt.Println("Add connection to station")
	fmt.Println(connections.Connections)
	stations := models.GetStationsMap()
	// var tempStation models.Station
	for _, connection := range connections.Connections {
		// check all connections
		for _, station := range stations.StationsMap {
			// checks all station
			if station.Name == connection.StationOne {
				// if the station matches find the corresponding station, and add to the structure
				// stations.StationsMap[i].Connections = append(stations.StationsMap[i].Connections, station) //TODO need to find and update not add
				
				models.StationsInstance.UpdateStationConnection(connection)//TODO need to find and update not add
			}
		}
	}
	for _, station := range stations.StationsMap {
		fmt.Println(station.Name)
		fmt.Println(station.Connections)
	}
}
