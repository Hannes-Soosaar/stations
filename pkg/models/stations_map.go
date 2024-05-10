package models

import (
	"fmt"
	"sync"
)

type StationsMap struct {
	StationsMap []Station
}

var StationsInstance *StationsMap
var StationsOnce sync.Once

func GetStationsMap() *StationsMap {
	StationsOnce.Do(func() {
		StationsInstance = &StationsMap{}
	})
	return StationsInstance
}

func (s *StationsMap) UpdateStation(stationToUpdate Station) error {
	for i, station := range s.StationsMap {
		if station.Name == stationToUpdate.Name {
			// Update the station at index i
			s.StationsMap[i] = stationToUpdate
			return nil
		}
	}
	return fmt.Errorf("station with name %s not found", stationToUpdate.Name)
}
func (s *StationsMap) UpdateStationConnection(connectionToUpdate Connection) error {
	fmt.Println("updating connection")

	for _, station := range s.StationsMap {
		for _, connection := range station.ConnObj {
			fmt.Println("Station ONE:")
			fmt.Println(connection.StationOne)
			fmt.Println(connection.StationTwo)
			fmt.Println("Station TWO:")
			fmt.Println(connectionToUpdate.StationOne)

			fmt.Println(connectionToUpdate.StationTwo)
			if connection.StationOne == connectionToUpdate.StationOne && connection.StationTwo == connectionToUpdate.StationTwo {
				connection.Distance = connectionToUpdate.Distance
			}
		}
	}
	return fmt.Errorf("station with name %s not found", connectionToUpdate.StationOne)
}
