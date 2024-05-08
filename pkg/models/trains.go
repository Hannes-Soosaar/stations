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

func (s *Trains) AddTrainStation(trainId int ,TrainAt Station,){
	var tempTrain Train
	tempTrain.Id = trainId
	tempTrain.Location= TrainAt
	s.Trains = append(s.Trains, tempTrain )
}

func (s *Trains) UpdateTrainStation(trainId int ,TrainAt Station,){
	fmt.Printf("Adding Train to %d \n ", trainId)
	fmt.Printf("the lendght of the string of trains is: %d \n",len(s.Trains))
	var tempTrain Train
	tempTrain.Id = trainId
	tempTrain.Location= TrainAt
	s.Trains = append(s.Trains, tempTrain )


	// for i, _ := range s.Trains {
	// 		s.Trains[i].Id = trainId
	// 		s.Trains[i].Location = TrainAt
	// }
}
