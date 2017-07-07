package export

import (
	"proto"
	"fmt"
)

type MySceneManager struct {
	manager 		GameManager
}

func NewSceneManager() *MySceneManager {
	return &MySceneManager{
	}
}

func (sm *MySceneManager) RegisterScene(opt *SceneOption) error {
	return nil
}

func (sm *MySceneManager) KickClient(id uint32, reason *proto.Message) {
	sm.manager.KickClient(id, reason)
}

func (sm *MySceneManager) SendClientMessage(id uint32, m *proto.Message) {
	sm.manager.SendClientMessage(id, m)
}

func (sm *MySceneManager) BroadcastMessage(ids []uint32, m *proto.Message) {
	sm.manager.BroadcastMessage(ids, m)
}

func (sm *MySceneManager) SendCommandInfo(source uint32, m *proto.Message) {
	sm.manager.SendCommandInfo(source, m)
}

func (sm *MySceneManager) OnInit(mgr GameManager) {
	sm.manager = mgr
}

func (sm *MySceneManager) OnClientConnected(id uint32, addr string) {
	fmt.Println(id, addr)
}

func (sm *MySceneManager) OnClientDisconnected(id uint32) {

}

func (sm *MySceneManager) OnClientMessage(id uint32, m *proto.Message) {

}

func (sm *MySceneManager) OnCommandInfo(source uint32, m *proto.Message) {

}

