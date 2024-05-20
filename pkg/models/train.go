package models

type Train struct {
	Id              int
	CurrentStation  string // tracks the station where the train is at, is set to start.
	LastStation     string // not using
	NextStation     string // not using
	TrainOnRout     int
	IsAtDestination bool // End station will switch to  false
	// if the LocationName == Instance.EndStation  removes the train  this is OK
}
