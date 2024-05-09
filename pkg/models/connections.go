package models

import (
	"fmt"
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
	for i, connection := range s.Connections {
		if connection.StationOne == c.StationOne || connection.StationTwo == c.StationTwo {
			s.Connections[i] = c
			return nil
		}
	}
	return fmt.Errorf("station with station name  %s not found", c.StationOne)
}
// TODO add method to update the connections ?
