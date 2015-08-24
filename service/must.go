package service

import (
	"github.com/yangsf5/claw/center"
)


func Register() {
	services := map[string]center.Service{
		"Master": &Master{},
		"Harbor": &Harbor{},
		"Error": &Error{},
		"Test": &Test{},
		"Gate": &Gate{},
		"Web": &Web{},
	}

	for name, cb := range services {
		center.Register(name, cb)
	}
}

func send(source, destination string, session int, msgType int, msg interface{}) {
	center.Send(source, destination, session, msgType, msg)
}
