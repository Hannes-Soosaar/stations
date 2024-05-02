package models

import "sync"

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
