package utils

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func openMapFromFile(path string) {
	var stations models.StationsMap
	fmt.Println("OPENING FROM INSTANCE " + path)
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
			addStationsToMap(line)
		} else if isConnectionSection {
			if line == "connections:" {
				continue
			}
			connections = append(connections, line)
		} else {
			fmt.Println("not in any section!")
		}

	}
	mapConnections(connections)
	getConnections(stations)
	instance := models.GetInstance()
	createTrains(instance.NumberOfTrains, instance.StartStation)
}

func createTrains(NumberOfTrains int, StartStation string) models.Trains {
	var trains models.Trains
	stations := models.GetStationsMap()
	var TrainLocation models.Station
	for _, station := range stations.StationsMap {
		if StartStation == station.Name {
			TrainLocation = station
			break
		}
	}

	for i := 0; i < NumberOfTrains; i++ {
		train := models.Train{
			Id:       i,
			Location: TrainLocation,
		}
		trains.Trains = append(trains.Trains, train)
	}
	fmt.Println("Trains: ")
	fmt.Println(trains)
	return trains
}
