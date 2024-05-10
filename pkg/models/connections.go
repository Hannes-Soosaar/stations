package models

import (
	"fmt"
	"log"
	"sync"
)

type Connections struct {
	Connections []Connection
}

var connectionsInstance *Connections
var connectionsOnce sync.Once

func GetConnectionsP() *Connections {
	connectionsOnce.Do(func() {
		connectionsInstance = &Connections{}
	})
	return connectionsInstance
}

func (s *Connections) UpdateConnections(c Connection) error {
	// fmt.Println("Updating Connection %d", c.Distance)
		// fmt.Println("Getting connections")
		// fmt.Println(c)
	for i, connection := range s.Connections {
		if connection.StationOne == c.StationOne && connection.StationTwo == c.StationTwo {
			log.Println("updating")
			log.Println(s.Connections) 
			log.Println(c)
			s.Connections[i].Distance = c.Distance
			return nil
		}else{
			s.Connections[i] = s.Connections[i]
		}
	}
	return fmt.Errorf("station with station name  %s not found", c.StationOne)
}
// TODO add method to update the connections ?
