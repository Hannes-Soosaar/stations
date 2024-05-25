package utils

import (
	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

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
}

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
