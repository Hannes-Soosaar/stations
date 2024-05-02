package models


//TODO create a slice of trains based on the number of trains in the flag perhaps its important to have the 
type Train struct {
	Id int
	Location Station // tracks the station where the train is at, is set to start. 
}
