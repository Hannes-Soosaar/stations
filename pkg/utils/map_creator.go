package utils

import (
	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// get connectionsist

func addConnectionToStations() {
	connections := models.GetConnectionsP()
	stations := models.GetStationsMap()
	for _, connection := range connections.Connections {
		// check all connections
		for i, station := range stations.StationsMap {
			// checks all station
			if station.Name == connection.StationOne {
				// if the station matches find the corresponding station, and add to the structure
				stations.StationsMap[i].Connections = append(stations.StationsMap[i].Connections)
			}
		}
	}
}
