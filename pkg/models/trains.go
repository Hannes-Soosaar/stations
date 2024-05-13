package models

import (
	// "fmt"
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

func (s *Trains) AddTrainStation(trainId int, TrainAt Station) {
	var tempTrain Train
	tempTrain.Id = trainId
	tempTrain.LocationName = TrainAt.Name
	s.Trains = append(s.Trains, tempTrain)
}

//	func (s *Trains) UpdateTrainStation(trainId int, TrainAt string) {
//		var tempTrain Train
//		tempTrain.Id = trainId
//		tempTrain.LocationName = TrainAt
//		s.Trains = append(s.Trains, tempTrain)
//	}
func (s *Trains) UpdateTrainLocation(trainId int, TrainAt string) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].LocationName = TrainAt
		}
	}
}
