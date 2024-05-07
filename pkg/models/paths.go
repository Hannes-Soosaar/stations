package models

import "sync"

type Path struct {
	PathStations  []Station
	TrainLocation Station
}

type Paths struct {
	Paths []Path
}

var instancePaths *Paths
var oncePath sync.Once

func GetPaths() *Paths {

	oncePath.Do(func() {
		instancePaths = &Paths{}
	})
	return instancePaths
}

func (paths *Paths) AddPath(path Path) {
	paths.Paths = append(paths.Paths, path)
}
