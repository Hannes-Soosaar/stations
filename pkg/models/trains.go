package models

import (
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
	tempTrain.CurrentStation = TrainAt.Name
	s.Trains = append(s.Trains, tempTrain)
}

func (s *Trains) RemoveTrainById(trainId int) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains = append(s.Trains[:i], s.Trains[i+1:]...)
		}
	}

}
func (s *Trains) SetArrivedAtDestinationById(trainId int) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].IsAtDestination = true 
		}
	}
}
func (s *Trains) SetDestinationPrintedById(trainId int) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].DestinationPrinted = true 
		}
	}
}

func (s *Trains) UpdateTrainLocation(trainId int, TrainAt string) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].CurrentStation = TrainAt
		}
	}
}
func (s *Trains) UpdateTrainOnRout(trainId int, routNumber int) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].TrainOnRout = routNumber
		}
	}
}
func (s *Trains) UpdateTrainNextLocation(trainId int, TrainAt string) {
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].NextStation = TrainAt
		}
	}
}
