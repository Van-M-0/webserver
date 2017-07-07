package gameserver

import "proto"

type GameDispatcher struct {

}

func (gf *GameFrame) OnClientConnected(id uint32, addr string) {

}

func (gf *GameFrame) OnClientDisconnected(id uint32) {

}

func (gf *GameFrame) OnClientMessage(id uint32, m *proto.Message) {

}

func (gf *GameFrame) OnCommandInfo(source uint32, m *proto.Message) {

}
