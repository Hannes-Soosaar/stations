package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// return an error.
func GetAndCheckInput() (string, string, string, int, error) {

	var errorMessage string

	if len(os.Args) < 5 {
		errorMessage = "too few command line arguments. Correct usage is go run . [path to file containing network map] [start station] [end station] [number of trains]"
		return "", "", "", 0, fmt.Errorf(errorMessage)
	}

	if len(os.Args) > 5 {
		errorMessage = "too many command line arguments. Correct usage is go run . [path to file containing network map] [start station] [end station] [number of trains]"
		return "", "", "", 0, fmt.Errorf(errorMessage)
	}

	networkMap := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	trainAmountStr := os.Args[4]

	dir := "../assets/input/" // will be messed up if somebody wants to add the full path. eg. c:/myMap

	startStationFound := false
	endStationFound := false

	mapFile, err := os.Open(dir + networkMap)

	if err != nil {
		return "", "", "", 0, fmt.Errorf("error opening network map file: %v", err)
	}

	defer mapFile.Close()
	scanner := bufio.NewScanner(mapFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, startStation) {
			startStationFound = true
		}
		if strings.HasPrefix(line, endStation) {
			endStationFound = true
		}
	}

	// File operation error
	if err := scanner.Err(); err != nil {
		return "", "", "", 0, fmt.Errorf("error scanning network map file: %v", err)
	}

	// Map validation
	if !startStationFound {
		errorMessage += "Entered starting station does not exist in this map. "
	}
	// Map validation
	if !endStationFound {
		errorMessage += "Entered ending station does not exist in this map. "
	}
	// input CLI
	if startStation == endStation {
		errorMessage += "The start and end stations cannot be the same. "
	}
	// input CLI
	trainAmount, err := strconv.Atoi(trainAmountStr)
	if err != nil {
		errorMessage += "Train amount has to be a number. "
	} else {
		if trainAmount < 1 {
			errorMessage += "There has to be at least 1 train. "
		}
	}
	// return all CLI errors.
	if errorMessage != "" {
		return "", "", "", 0, fmt.Errorf(errorMessage)
	}

	// check CLI input, if OK create instance.
	instance := models.Instance{
		PathToMap:      os.Args[1],
		StartStation:   os.Args[2],
		EndStation:     os.Args[3],
		NumberOfTrains: os.Args[4],
	}

	// create instance stationsMap
	openMapFromFile(instance.PathToMap)

	
	return networkMap, startStation, endStation, trainAmount, nil
}
