package utils

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func openMapFromFile(path string){
	var stations []models.Station
	fmt.Println("OPENING FROM INSTANCE " + path)

	// stationsMap := models.GetStationsMap() // saves the pointer

	var connections []string

	mapFile, err := os.Open("../assets/input/" + path) // move the input directory to config
	if err != nil {
		fmt.Errorf("error opening network map file: %v", err) // move error text to internal error_codes
	}

	defer mapFile.Close()
	isStationSection := false
	isConnectionSection := false

	scanner := bufio.NewScanner(mapFile)

	for scanner.Scan() {

		line := scanner.Text()
		var clearLine string
		for _, r := range line {
			if r == '#' {
				line = clearLine
				break
			} else if unicode.IsSpace(r) {
				continue
			} else {
				clearLine += string(r)
			}
			line = clearLine
		}

		if line == "stations:" {
			isStationSection = true
			isConnectionSection = false
		} else if line == "connections:" {
			isConnectionSection = true
			isStationSection = false
		} else if line == "" {
			continue
		}

		if isStationSection {
			if line == "stations:" {
				continue
			}
			station := getStation(line)
			stations = append(stations, station)
		} else if isConnectionSection {
			if line == "connections:" {
				continue
			}
			connections = append(connections, line) // Create a slice with all stations 
		} else {
			fmt.Println("not in any section!")
		}

	}

	// mapStations(stations)
	mapConnections(connections)
	// getConnections(stationsMap.StationsMap, connections) // Broke this for a bit
}

// func handleStuff() {
// 	if err := scanner.Err(); err != nil {
// 		return "", "", "", 0, fmt.Errorf("error scanning network map file: %v", err)
// 	}

// 	if !startStationFound {
// 		errorMessage += "Entered starting station does not exist in this map. "
// 	}

// 	if !endStationFound {
// 		errorMessage += "Entered ending station does not exist in this map. "
// 	}

// 	if startStation == endStation {
// 		errorMessage += "The start and end stations cannot be the same. "
// 	}

// 	trainAmount, err := strconv.Atoi(trainAmountStr)
// 	if err != nil {
// 		errorMessage += "Train amount has to be a number. "
// 	} else {
// 		if trainAmount < 1 {
// 			errorMessage += "There has to be at least 1 train. "
// 		}
// 	}

// 	if errorMessage != "" {
// 		return "", "", "", 0, fmt.Errorf(errorMessage)
// 	}
// 	return networkMap, startStation, endStation, trainAmount, nil
