package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// TODO this code should be refactored

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
	startStationFound := false
	endStationFound := false
	dir := ""

	if !strings.Contains(networkMap, "../assets/tests/input") {
		dir = "../assets/input/"
		networkMap = filepath.Join(dir, networkMap)
	}

	mapFile, err := os.Open(networkMap)
	fmt.Println(networkMap)
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
	if err := scanner.Err(); err != nil {
		return "", "", "", 0, fmt.Errorf("error scanning network map file: %v", err)
	}
	if !startStationFound {
		errorMessage += "Error the entered START station " + startStation + " does not exist on the:" + networkMap
	}
	if !endStationFound {
		errorMessage += "Error the entered END station " + endStation + " does not exist on the:" + networkMap
	}
	if startStation == endStation {
		errorMessage += "The START: " + startStation + " and END: " + endStation + " stations cannot be the same. "
	}
	trainAmount, err := strconv.Atoi(trainAmountStr)
	if err != nil {
		errorMessage += "error: train amount has to be a number. "
	} else {
		if trainAmount == 0 {
			errorMessage += "error: there has to be at least 1 train. "
		} else if trainAmount < 0 {
			errorMessage += " error: the number of trains has to be a positive integer. "
		}
	}
	if errorMessage != "" {
		return "", "", "", 0, fmt.Errorf(errorMessage)
	}

	models.InitInstance(networkMap, startStation, endStation, trainAmount)
	instance := models.GetInstance()
	// create instance stationsMap
	openMapFromFile(instance.PathToMap)
	return networkMap, startStation, endStation, trainAmount, nil
}
