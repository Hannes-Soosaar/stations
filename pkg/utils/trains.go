package utils

import (
	"log"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// ! Works
func createTrains() {
	instance := models.GetInstance()
	trains := models.GetTrains()
	for i := 0; i < instance.NumberOfTrains; i++ {
		train := models.Train{
			Id:             i,
			CurrentStation: instance.StartStation,
		}
		trains.Trains = append(trains.Trains, train)
	}
	log.Println(trains.Trains)
}

// // !Works
// func MoveTrains() {
// 	trains := models.GetTrains()
// 	for _, train := range trains.Trains {
// 		// GetNextStation(train.Id)
// 	}
// }

func findCurrentStationName(trainId int) string {
	trains := models.GetTrains()
	currentStation := ""
	for _, train := range trains.Trains {
		if train.Id == trainId {
			currentStation = train.CurrentStation
		}
	}
	return currentStation
}

func findTrainById(trainId int) models.Train {
	trains := models.GetTrains()
	var currentTrain models.Train
	for _, train := range trains.Trains {
		if train.Id == trainId {
			currentTrain = train
		}
	}
	return currentTrain
}
func findLastStationName(trainId int) string {
	trains := models.GetTrains()
	lastStation := ""
	for _, train := range trains.Trains {
		if train.Id == trainId {
			lastStation = train.LastStation
		}
	}
	return lastStation
}

func updatedTrainCurrentStation(trainId int, newStation string) {
	// add function to update the location to the new station
	// same time it needs to set the current station to the last station.
}
