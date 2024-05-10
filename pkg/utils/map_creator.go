package utils

import (
	// "fmt"

	"log"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addConnectionToStations() {
	connections := models.GetConnectionsP()
	stations := models.GetStationsMap()
	// var tempStation models.Station
	for _, connection := range connections.Connections {
		// check all connections
		for _, station := range stations.StationsMap {
			// checks all station
			if station.Name == connection.StationOne {
				models.StationsInstance.UpdateStationConnection(connection) //TODO need to find and update not add
			}
		}
	}
	for _, station := range stations.StationsMap {
		log.Println(station.Name)
		log.Println(station.ConnObj)
	}

}
