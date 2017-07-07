package lobby

import (
	"proto"
	"scenemgr"
)

type LobbyManager struct {
	lsServer 		*LobbServer
	manager 		scenemgr.SceneManager
}

func (gm *LobbyManager) KickClient(id uint32, reason *proto.Message) {
	gm.lsServer.gwServer.KickClient(id)
}

func (gm *LobbyManager) SendClientMessage(id uint32, m *proto.Message) {
	gm.lsServer.gwServer.SendMessage(id, m)
}

func (gm *LobbyManager) BroadcastMessage(ids []uint32, m *proto.Message) {
	gm.lsServer.gwServer.BroadcastMessage(ids, m)
}

func (gm *LobbyManager) SendCommandInfo(source uint32, m *proto.Message) {

}

func(gm *LobbyManager) OnClientConnected(id uint32, addr string) {
	gm.manager.OnClientConnected(id, addr)
}

func(gm *LobbyManager) OnClientDisconnected(id uint32) {
	gm.manager.OnClientDisconnected(id)
}

func(gm *LobbyManager) OnClientMessage(id uint32, m *proto.Message) {
	gm.manager.OnClientMessage(id, m)
}

func(gm *LobbyManager) OnCommandInfo(source uint32, m *proto.Message) {
	gm.manager.OnCommandInfo(source, m)
}
