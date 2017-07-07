package gameserver

import (
	"proto"
)

type GameFrame struct {
	gameServer 	*GameServer
}

func (gf *GameFrame) KickClient(id uint32, reason *proto.Message) {

}

func (gf *GameFrame) SendClientMessage(id uint32, m *proto.Message) {

}

func (gf *GameFrame) BroadcastMessage(ids []uint32, m *proto.Message) {

}


func (gf *GameFrame) SendCommandInfo(source uint32, m *proto.Message) {

}

