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