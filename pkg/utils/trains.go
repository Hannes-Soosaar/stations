package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func createTrains(){
	instance := models.GetInstance()
	// trains := models.GetTrains() 
	

	stations := models.GetStationsMap()
	var TrainLocation models.Station
	for _, station := range stations.StationsMap {
		if instance.StartStation == station.Name {
			TrainLocation = station
			break
		}
	}

	for i := 0; i < instance.NumberOfTrains; i++ {
		models.GetTrains().AddTrainStation(i, TrainLocation)
	// 	train := models.Train{
	// 		Id:       i,
	// 		Location: TrainLocation,
	// 	}
	// 	trains.Trains = append(trains.Trains, train)
	}

	// fmt.Println("Instance")
	// for _,train := range trains.Trains{
	// 	fmt.Println(train.Location.Name)
	// }
	// return trains
}

func MoveTrains(){

	trains:= models.GetTrains()
	fmt.Println("HERE")
	fmt.Println(trains.Trains)
	for _,train := range trains.Trains{
		fmt.Println("Moving Trains!")
		GetShortestPath(train.Id)
		//TODO add a function to move the train
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



