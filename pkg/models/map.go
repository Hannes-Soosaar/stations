package models

// stores all the nodes(stations) all the edges to the node in station.connection
// we remove all station that have been visited after each pass.
type Map struct{
	StationsMap []Station
}