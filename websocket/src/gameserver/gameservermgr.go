package gameserver

import (
	"proto"
	"scenemgr"
)

type GameServerMgr struct {
	gameServer 	*GameServer
	manager 	scenemgr.SceneManager
}

func (gm *GameServerMgr) KickClient(id uint32, reason *proto.Message) {
	gm.gameServer.gwServer.KickClient(id)
}

func (gm *GameServerMgr) SendClientMessage(id uint32, m *proto.Message) {
	gm.gameServer.gwServer.SendMessage(id, m)
}

func (gm *GameServerMgr) BroadcastMessage(ids []uint32, m *proto.Message) {
	gm.gameServer.gwServer.BroadcastMessage(ids, m)
}

func (gm *GameServerMgr) SendCommandInfo(source uint32, m *proto.Message) {

}

func(gm *GameServerMgr) OnClientConnected(id uint32, addr string) {
	gm.manager.OnClientConnected(id, addr)
}

func(gm *GameServerMgr) OnClientDisconnected(id uint32) {
	gm.manager.OnClientDisconnected(id)
}

func(gm *GameServerMgr) OnClientMessage(id uint32, m *proto.Message) {
	gm.manager.OnClientMessage(id, m)
}

func(gm *GameServerMgr) OnCommandInfo(source uint32, m *proto.Message) {
	gm.manager.OnCommandInfo(source, m)
}

