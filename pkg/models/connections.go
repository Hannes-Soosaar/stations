package models

import "sync"

type Connections struct{

	Connections []Connection

}

var instance *Connections 
var once sync.Once

func GetConnectionsP() *Connections {

once.Do(func(){
	instance = &Connections{}
})
	return instance
}

