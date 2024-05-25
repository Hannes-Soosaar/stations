package models

type Station struct {
	Name        string  // filled in from the file
	X           float64 // filled in from the file
	Y           float64 //	filled in from the file
	IsVisited   bool    // set to null
	IsOccupied  bool    // set to null
	IsStart     bool    //
	IsFinish    bool
	IsMapped    bool         //  not sure if we will need this
	Connections []Station    // is filled in from the file. holds the information on how many edges there are.
	ConnObj     []Connection // Try to us this instead of the slice of stations.
}

func (s *Station) RemoveConnection(name string) {
	for i, conn := range s.Connections {
		if conn.Name == name {
			s.Connections = append(s.Connections[:i], s.Connections[i+1:]...)
			return
		}
	}
}
