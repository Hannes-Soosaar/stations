package main

import (
	"fmt"
	"stations/pkg/utils"
)

func main() {
	networkMap, startStation, endStation, trainAmount, err := utils.GetAndCheckInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Network Map:", networkMap)
	fmt.Println("Start Station:", startStation)
	fmt.Println("End Station:", endStation)
	fmt.Println("Train Amount:", trainAmount)
}
