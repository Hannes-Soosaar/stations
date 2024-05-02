package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func PrintToConsol(){
	connections:= models.GetConnectionsP()
	stations := models.GetStationsMap()
	fmt.Println("printing singletons")
	fmt.Println(stations)
	fmt.Println(connections)
}