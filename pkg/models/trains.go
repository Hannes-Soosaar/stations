package models

import "sync"

type Trains struct {
	Trains []Train
}

var instanceT *Trains
var onceT sync.Once

func GetTrains() *Trains {

	onceT.Do(func() {
		instanceT = &Trains{}
	})
	return instanceT
}
