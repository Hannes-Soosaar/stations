package utils

import (
	"fmt"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func getConnections(stationsMap models.StationsMap, connections []string) models.StationsMap {
	var stationNames []string
	var splitConnections [][]string
	var stationToAppend string = ""

	for _, station := range stationsMap.StationsMap {
		stationNames = append(stationNames, station.Name)
	}

	for _, connection := range connections {
		splitConnection := strings.Split(connection, "-")
		splitConnections = append(splitConnections, splitConnection)
	}

	for i, stationName := range stationNames {
		for _, splitConnection := range splitConnections {
			if stationName == splitConnection[1] {
				stationToAppend = splitConnection[0]

			} else if stationName == splitConnection[0] {
				stationToAppend = splitConnection[1]
			}

			if stationToAppend != "" {
				for _, currentStation := range stationsMap.StationsMap {
					if currentStation.Name == stationToAppend {
						stationsMap.StationsMap[i].Connections = append(stationsMap.StationsMap[i].Connections, currentStation)
						stationToAppend = ""
					}
				}

			}
		}
	}

	//testing if everything works
	fmt.Println("")
	for _, station := range stationsMap.StationsMap {
		fmt.Printf("Station name: %s\n", station.Name)
		fmt.Printf("X coordinate: %v\n", station.X)
		fmt.Printf("Y coordinate: %v\n", station.Y)
		fmt.Printf("Connected stations:\n")
		for _, connection := range station.Connections {
			fmt.Printf("- %s\n", connection.Name)
		}
		fmt.Println("----------------------")
	}
	return stationsMap
}
