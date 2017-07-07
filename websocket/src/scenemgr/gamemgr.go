package scenemgr

import "proto"

type SceneOption struct {

}

type GameManager interface {
	KickClient(id uint32, reason *proto.Message)
	SendClientMessage(id uint32, m *proto.Message)
	BroadcastMessage(ids []uint32, m *proto.Message)
	SendCommandInfo(source uint32, m *proto.Message)
}

type SceneManager interface {
	OnInit(manager GameManager)
	OnClientConnected(id uint32, addr string)
	OnClientDisconnected(id uint32)
	OnClientMessage(id uint32, m *proto.Message)
	OnCommandInfo(source uint32, m *proto.Message)
	KickClient(id uint32, reason *proto.Message)
	SendClientMessage(id uint32, m *proto.Message)
	BroadcastMessage(ids []uint32, m *proto.Message)
	SendCommandInfo(source uint32, m *proto.Message)
	CreateScene(opt *SceneOption) Scene
}

type Scene interface {
	OnCreated(mgr *SceneManager)
	OnDestroyed()
}




