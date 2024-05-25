package models

import (
	"errors"
	"fmt"

	// "log"
	"sync"
)

type Connections struct {
	Connections []Connection
}

var connectionsInstance *Connections
var connectionsOnce sync.Once

func GetConnectionsP() (*Connections,error) {
	var err error
	connectionsOnce.Do(func() {
		connectionsInstance = &Connections{}
	if connectionsInstance == nil { 
			err = errors.New("there are no connections between stations")
		}
	})
	if err != nil {
		return nil,err
	}
	return connectionsInstance, err
}

func (s *Connections) UpdateConnections(c Connection) error {
	fmt.Println(s.Connections)
	for i, connection := range s.Connections {
		if connection.StationOne == c.StationOne && connection.StationTwo == c.StationTwo {
			s.Connections[i].Distance = c.Distance
			return nil
		}
	}
	return fmt.Errorf("station with station name  %s not found", c.StationOne)
}
