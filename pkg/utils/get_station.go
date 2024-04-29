package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/stations/pkg/models"
)

func getStation(line string) models.Station {

	newStation := models.Station{}
	params := strings.Split(line, ",")

	newStation.Name = params[0]

	x, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	y, err := strconv.Atoi(params[2])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	newStation.X = x
	newStation.Y = y

	return newStation
}
