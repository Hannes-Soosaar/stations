package utils
// the file is named wrong. It gets connections not connection
import (
	"fmt"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func mapConnections(cs []string) {
	fmt.Println(cs)

	var connection models.Connection 
	connections := models.GetConnectionsP();

	for _,c :=range cs {
		// split the connection 
		split := strings.Split(c, "-")
		if len(split) == 2 {  // does not add anything with more than one station
				connection.StationOne = split[0]
				connection.StationTwo = split[1]
		} else {
			fmt.Println("not a valid connection")
		}
		connections.Connections= append(connections.Connections,connection)
	}
	fmt.Println(connections.Connections)
}


// The function is named wrong it creates a StationsMap. 
func getConnections(stationsMap models.StationsMap) models.StationsMap {
	var stationNames []string
	// var splitConnections [][]string // already done by mapConnections function.
	allConnections := models.GetConnectionsP()
	var stationToAppend string = ""

	for _, station := range stationsMap.StationsMap {
		stationNames = append(stationNames, station.Name)
	}

	// for _, connection := range connections {
	// 	splitConnection := strings.Split(connection, "-")
	// 	splitConnections = append(splitConnections, splitConnection)
	// }
//TODO: do we want to iterate through all stations, or all connections ?
	for i, stationName := range stationNames {
		// fmt.Println(len(splitConnections))
// Loop over allConnections. OK

		// for _, splitConnection := range splitConnections {
		// 	if stationName == splitConnection[1] {
		// 		stationToAppend = splitConnection[0] 
		// 	} else if stationName == splitConnection[0] {
		// 		stationToAppend = splitConnection[1]
		// 	}

		for _, connection := range allConnections.Connections {
			fmt.Println(connection)
			if stationName == connection.StationOne {
				stationToAppend = connection.StationTwo
			} else if stationName == connection.StationTwo {
				stationToAppend = connection.StationOne
			}
			fmt.Println(stationToAppend)
// new Function this is the get station by name..
			if stationToAppend != "" {
				for _, currentStation := range stationsMap.StationsMap {
					if currentStation.Name == stationToAppend {
						stationsMap.StationsMap[i].Connections = append(stationsMap.StationsMap[i].Connections, currentStation)
						stationToAppend = ""
					}
// findStationByName(stationToAppend)
// new function add connection to station
//TODO: add the connection to the station returned from findStationByName 
//TODO: Do we need to update the stationMap struct?
					stationsMap.StationsMap[i].Connections = append(stationsMap.StationsMap[i].Connections, currentStation)
				}
			}
		}
	}

	// fmt.Println("")
	// for _, station := range stationsMap.StationsMap {
	// 	fmt.Printf("Station name: %s\n", station.Name)
	// 	fmt.Printf("X coordinate: %v\n", station.X)
	// 	fmt.Printf("Y coordinate: %v\n", station.Y)
	// 	fmt.Printf("Connected stations:\n")
	// 	for _, connection := range station.Connections {
	// 		fmt.Printf("- %s\n", connection.Name)
	// 	}
	// 	fmt.Println("----------------------")
	// }
	return stationsMap
}
