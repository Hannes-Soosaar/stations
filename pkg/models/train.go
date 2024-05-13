package models

type Train struct {
	Id           int
	LocationName string // tracks the station where the train is at, is set to start.
	LastStation  string
}
