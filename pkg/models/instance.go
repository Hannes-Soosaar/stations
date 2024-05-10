package models

type Instance struct {
	PathToMap      string
	StartStation   string
	EndStation     string
	NumberOfTrains int
}

var inputInstance *Instance

func InitInstance(path string, startStation string, endStation string, trainAmount int) {
	if inputInstance == nil {
		inputInstance = &Instance{
			PathToMap:      path,
			StartStation:   startStation,
			EndStation:     endStation,
			NumberOfTrains: trainAmount,
		}
	}
}

func GetInstance() *Instance {
	return inputInstance
}
