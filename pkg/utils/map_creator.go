package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func addConnectionToStations() {
	connections,err := models.GetConnectionsP()
	if err !=nil {
		fmt.Println(err)
	}

	if len(connections.Connections)>1{
		fmt.Errorf("There are no connections")
	}
	
	stations := models.GetStationsMap()


	for _, connection := range connections.Connections {
		for _, station := range stations.StationsMap {
			if station.Name == connection.StationOne {
				models.StationsInstance.UpdateStationConnection(connection)
			}
		}
	}
}
