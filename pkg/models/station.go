package models

// a node
// read in the station from file.
// create a connections slice with all connected stations
// need to calculate the distance?

type Station struct {
	Name        string // filled in from the file
	X           int    // filled in from the file
	Y           int    //	filled in from the file
	IsVisited   bool   // set to null
	IsOccupied  bool   // set to null
	IsStart     bool   //
	IsFinish    bool
	IsMapped    bool      //  not sure if we will need this
	Connections []Station // is filled in from the file. holds the information on how many edges there are.
}
