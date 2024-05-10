package utils

import (
	// "log"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)
//! Works
func createTrains(){
	instance := models.GetInstance()
	trains := models.GetTrains() 
	stations := models.GetStationsMap()
	var TrainLocation models.Station
	for _, station := range stations.StationsMap {
		if instance.StartStation == station.Name {
			TrainLocation = station
			break
		}
	}
	for i := 0; i < instance.NumberOfTrains; i++ {
		train := models.Train{
			Id:       i,
			Location: TrainLocation,
		}
		// log.Println("Adding train ")
		// log.Println(train.Id)
		trains.Trains = append(trains.Trains, train)
	}
}

//!Works
func MoveTrains(){
	trains:= models.GetTrains()
	for _,train := range trains.Trains{
	// log.Println(train.Location.Name)
		GetShortestPath(train.Id)
	}
}

func findCurrentStationName(trainId int) string {
	trains := models.GetTrains()
	currentStation :=""
	for _, train := range trains.Trains{
		if train.Id == trainId {
			currentStation = train.Location.Name
		}
	}
 	return currentStation
}

func updatedTrainCurrentStation(trainId int){

}



