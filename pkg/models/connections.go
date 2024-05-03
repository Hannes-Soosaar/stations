package models

import "sync"

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
