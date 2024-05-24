package models

import (
	"fmt"
	// "log"
	"sync"
)

type Connections struct {
	Connections []Connection
}

var connectionsInstance *Connections
var connectionsOnce sync.Once
//TODO: needs to give out an error if the connections is called but it remains empty!
func GetConnectionsP() *Connections {
	connectionsOnce.Do(func() {
		connectionsInstance = &Connections{}
	})
	return connectionsInstance
}

func (s *Connections) UpdateConnections(c Connection) error {
	for i, connection := range s.Connections {
		if connection.StationOne == c.StationOne && connection.StationTwo == c.StationTwo {
			s.Connections[i].Distance = c.Distance
			return nil
		}
	}
	return fmt.Errorf("station with station name  %s not found", c.StationOne)
}
