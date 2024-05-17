package models

import (
	"fmt"
	"sync"
	// "golang.org/x/text/unicode/rangetable"
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
			fmt.Println(trainId, "Removed")
		}
	}
}

func (s *Trains) UpdateTrainLocation(trainId int, TrainAt string) {
	// fmt.Printf("T %d %s \n", trainId, TrainAt)
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
			fmt.Printf("T %d %d \n", trainId, routNumber)
		}
	}
}
func (s *Trains) UpdateTrainNextLocation(trainId int, TrainAt string) {
	fmt.Printf("T%d, %s \n", trainId, TrainAt)
	for i, train := range s.Trains {
		if train.Id == trainId {
			s.Trains[i].NextStation = TrainAt
		}
	}
}
