package models

import (
	"fmt"
	"sync"
)

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

func (s *Trains) UpdateTrainStation(TrainAt Station,Id int) error {
	for i, train := range s.Trains {
		if train.Id == Id {
			s.Trains[i].Location = TrainAt
			return nil
		}
	}
	return fmt.Errorf("station with id %d not found", Id)
}
