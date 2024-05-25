package utils

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func openMapFromFile(path string) {
	var connections []string
	mapFile, err := os.Open(path) // move the input directory to config
	if err != nil {
		fmt.Errorf("error opening network map file: %v", err) // move error text to internal error_codes
	}
	defer mapFile.Close()
	isStationSection := false
	isConnectionSection := false
	hasStationSection := false
	hasConnectionStation := false
	stationCounter := 0
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
			hasStationSection = true

		} else if line == "connections:" {
			isConnectionSection = true
			isStationSection = false
			hasConnectionStation = true
		} else if line == "" {
			continue
		}
		if isStationSection {
			if line == "stations:" {
				
				continue
			}
			stationCounter++
			addStationsToMap(line)
		} else if isConnectionSection {
			if line == "connections:" {
				continue
			}
			connections = append(connections, line)
		} else {
			// fmt.Println("not in any section!")
		}
	}

	if !hasConnectionStation {
		err := fmt.Errorf("error: the \" %s \" has no Connections: section", path)
		fmt.Println(err)
		os.Exit(1)
	} else if !hasStationSection {
		err := fmt.Errorf("error: the \" %s \" has no Station: section", path)
		fmt.Println(err)
		os.Exit(1)
	}  
	if stationCounter > 10000 {
		err := fmt.Errorf("error: the \" %s \" has %d Stations. A map can not have more than 10 000 stations", path, stationCounter)
		fmt.Println(err)
		os.Exit(1)
	}
	mapConnections(connections)
	getConnections()
	checkForDuplicateCoordinates()
}
