// Author: sheppard(ysf1026@gmail.com) 2014-03-06

package service

import (
	"net"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/gate"
)

type Gate struct {
}

func (s *Gate) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.Master recv type=%v msg=%v", msgType, msg)
	switch msgType {
	case center.MsgTypeText:
		if msg, ok := msg.([]byte); ok&&session==0 {
			glog.Info(msg)
			gate.Broadcast(msg)
		}
	}
}

func (s *Gate) ClawStart() {
	go s.Listen()
}

func (s *Gate) Listen() {
	addr := center.BaseConfig.Gate.ListenAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	glog.Infof("Service.Gate listening, addr=%s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			glog.Errorf("Service.Gate accept error, err=%s", err.Error())
			continue
		}
		glog.Info("Service.Gate new connection")

		gate.ConnHandle(conn)
	}
}

