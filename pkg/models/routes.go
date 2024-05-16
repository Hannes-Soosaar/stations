package models

import "sync"

// This struct is not a singleton and can be created and destroyed at will
type Rout struct {
	StationNames []string
}

// Singleton that holds all possible routs.
type Routs struct {
	Routs []Rout
}

var instanceRouts *Routs
var onceRouts sync.Once

func GetRouts() *Routs {
	onceRouts.Do(func() {
		instanceRouts = &Routs{}
	})
	return instanceRouts
}

func (routs *Routs) AddRoutToRouts(rout Rout) {
	routs.Routs = append(routs.Routs, rout)
}
