package utils

import (
	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addConnectionToStations() {
	connections := models.GetConnectionsP()
	stations := models.GetStationsMap()
	for _, connection := range connections.Connections {
		for _, station := range stations.StationsMap {
			if station.Name == connection.StationOne {
				models.StationsInstance.UpdateStationConnection(connection)
			}
		}
	}
}
