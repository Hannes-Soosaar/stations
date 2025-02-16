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
			s.StationsMap[i] = stationToUpdate
			return nil
		}
	}
	return fmt.Errorf("error: station with name %s not found", stationToUpdate.Name)
}

func (s *StationsMap) UpdateStationConnection(connectionToUpdate Connection) error {
	for i, station := range s.StationsMap {
		for j, connection := range station.ConnObj {
			if connection.StationOne == connectionToUpdate.StationOne && connection.StationTwo == connectionToUpdate.StationTwo {
				s.StationsMap[i].ConnObj[j].Distance = connectionToUpdate.Distance
			}
		}
	}
	return fmt.Errorf("error: station with name %s not found", connectionToUpdate.StationOne)
}
