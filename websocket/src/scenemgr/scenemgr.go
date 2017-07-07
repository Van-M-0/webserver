package scenemgr

import (
	"proto"
	"fmt"
)

type GameSceneMgr struct {
	manager 		GameManager
}

func NewSceneManager() *GameSceneMgr {
	return &GameSceneMgr{
	}
}

func (sm *GameSceneMgr) RegisterScene(opt *SceneOption) error {
	return nil
}

func (sm *GameSceneMgr) KickClient(id uint32, reason *proto.Message) {
	sm.manager.KickClient(id, reason)
}

func (sm *GameSceneMgr) SendClientMessage(id uint32, m *proto.Message) {
	sm.manager.SendClientMessage(id, m)
}

func (sm *GameSceneMgr) BroadcastMessage(ids []uint32, m *proto.Message) {
	sm.manager.BroadcastMessage(ids, m)
}

func (sm *GameSceneMgr) SendCommandInfo(source uint32, m *proto.Message) {
	sm.manager.SendCommandInfo(source, m)
}

func (sm *GameSceneMgr) OnInit(mgr GameManager) {
	sm.manager = mgr
}

func (sm *GameSceneMgr) OnClientConnected(id uint32, addr string) {
}

func (sm *GameSceneMgr) OnClientDisconnected(id uint32) {

}

func (sm *GameSceneMgr) OnClientMessage(id uint32, m *proto.Message) {
	fmt.Println("game scene manager recv message .. ", id, m)
	sm.SendClientMessage(id, &proto.Message{
		Cmd: 101,
		Msg: &proto.GuestAuthMessage {
			Account: "hello world",
		},
	})
}

func (sm *GameSceneMgr) OnCommandInfo(source uint32, m *proto.Message) {

}

func (sm *GameSceneMgr) CreateScene(opt *SceneOption) Scene {
	return nil
}

