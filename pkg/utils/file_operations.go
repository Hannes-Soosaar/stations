package utils

import (
	"bufio"
	"fmt"
	"os"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func openMapFromFile(path string) models.StationsMap {

	fmt.Println("OPENING FROM INSTANCE " + path)

	var stationsMap models.StationsMap

	mapFile, err := os.Open("../assets/input/" + path) // move the input directory to config
	if err != nil {
		fmt.Errorf("error opening network map file: %v", err) // move error text to internal error_codes
		return stationsMap
	}

	defer mapFile.Close()
	isStationSection := false
	isConnectionSection := false

	scanner := bufio.NewScanner(mapFile)

	for scanner.Scan() {

		//TODO handle empty line.

		line := scanner.Text()

		for i, r := range line {
			if r == '#' {
				cleanLine := line[:i]
				line = cleanLine
				break
			}
		}

		if line == "stations:" {
			isStationSection = true
			isConnectionSection = false
		} else if line == "connections:" {
			isConnectionSection = true
			isStationSection = false
		} else if line == "" {
			continue
		} else {
			fmt.Println("reading content")
		}

		if isStationSection {
			if line == "stations:" {
				continue
			}
			getStation(line)
		} else if isConnectionSection {
			if line == "connections:" {
				continue
			}
			//TODO run function to add the connection to the station in the map.
		} else {
			fmt.Println("not in any section!")
		}

	}

	return stationsMap
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
