package models

type Train struct {
	Id             int
	CurrentStation string // tracks the station where the train is at, is set to start.
	LastStation    string // tracks the last station where the train was.
	NextStation    string
	// if the LocationName == Instance.EndStation  removes the train  this is OK
}
