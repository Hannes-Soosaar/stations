package models

type Station struct {
	Name       string  // filled in from the file
	X          float64 // filled in from the file
	Y          float64 //	filled in from the file
	IsVisited  bool    // set to null
	IsOccupied bool    // set to null
	IsStart    bool    //
	//Next move ?
	IsFinish    bool
	IsMapped    bool      //  not sure if we will need this
	Connections []Station // is filled in from the file. holds the information on how many edges there are.
}
