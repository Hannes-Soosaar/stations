package utils

import (
	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

// ! Works
func createTrains() {
	instance := models.GetInstance()
	trains := models.GetTrains()
	stations := models.GetStationsMap()
	var TrainLocation string
	for _, station := range stations.StationsMap {
		if instance.StartStation == station.Name {
			TrainLocation = station.Name
			break
		}
	}
	for i := 0; i < instance.NumberOfTrains; i++ {
		train := models.Train{
			Id:           i,
			LocationName: TrainLocation,
		}
		trains.Trains = append(trains.Trains, train)
	}
}

// !Works
func MoveTrains() {
	trains := models.GetTrains()
	for _, train := range trains.Trains {
		GetShortestPath(train.Id)
	}
}

func findCurrentStationName(trainId int) string {
	trains := models.GetTrains()
	currentStation := ""
	for _, train := range trains.Trains {
		if train.Id == trainId {
			currentStation = train.LocationName
		}
	}
	return currentStation
}

func updatedTrainCurrentStation(trainId int) {

}
