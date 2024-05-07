package main

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/stations/pkg/utils"
)

func main() {
	_, _, _, _, err := utils.GetAndCheckInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	utils.PrintToConsol()
	// fmt.Println("Network Map:", networkMap)
	// fmt.Println("Start Station:", startStation)
	// fmt.Println("End Station:", endStation)
	// fmt.Println("Train Amount:", trainAmount)
	utils.FindPath()
}
