package models

import "sync"

type Rout struct {
	StationNames []string
}
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
