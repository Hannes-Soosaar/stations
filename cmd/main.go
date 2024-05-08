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

	// TODO Launch instance 
	utils.PrintToConsol()
	utils.FindPath()
}
