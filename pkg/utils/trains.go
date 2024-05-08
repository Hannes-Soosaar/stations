package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func createTrains() models.Trains {
	instance := models.GetInstance()

	var trains models.Trains
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
		trains.Trains = append(trains.Trains, train)
	}
	fmt.Println("Instance")
	for _,train := range trains.Trains{
		fmt.Println(train.Location.Name)
	}
	return trains
}