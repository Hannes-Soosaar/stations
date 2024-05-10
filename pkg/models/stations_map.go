package models

import (
	"fmt"
	"log"
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

// ! It is important to use the exact iteration number, if you do not, it will update a copy not the reference!
func (s *StationsMap) UpdateStationConnection(connectionToUpdate Connection) error {
	log.Println(connectionToUpdate)
	for i, station := range s.StationsMap {
		for j, connection := range station.ConnObj {
			fmt.Println(connection)
			if connection.StationOne == connectionToUpdate.StationOne && connection.StationTwo == connectionToUpdate.StationTwo {
				s.StationsMap[i].ConnObj[j].Distance = connectionToUpdate.Distance
			}
		}
	}
	return fmt.Errorf("station with name %s not found", connectionToUpdate.StationOne)
}
