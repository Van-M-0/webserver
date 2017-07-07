package scenemgr

import (
	"proto"
	"fmt"
	"dbproxy"
)

type RoomSceneMgr struct {
	manager 		GameManager
	dbProxy 		*dbproxy.DbProxy
}

func NewRoomSceneManager() *RoomSceneMgr {
	return &RoomSceneMgr{
		dbProxy: dbproxy.NewDbProxy(&dbproxy.DbOption{
			Host: "127.0.0.1:3306",
			User: "root",
			Pass: "1",
			Name: "mygame",
			ShowDetailLog: true,
			Singular: true,
		}),
	}
}

func (sm *RoomSceneMgr) RegisterScene(opt *SceneOption) error {
	return nil
}

func (sm *RoomSceneMgr) KickClient(id uint32, reason *proto.Message) {
	sm.manager.KickClient(id, reason)
}

func (sm *RoomSceneMgr) SendClientMessage(id uint32, m *proto.Message) {
	sm.manager.SendClientMessage(id, m)
}

func (sm *RoomSceneMgr) BroadcastMessage(ids []uint32, m *proto.Message) {
	sm.manager.BroadcastMessage(ids, m)
}

func (sm *RoomSceneMgr) SendCommandInfo(source uint32, m *proto.Message) {
	sm.manager.SendCommandInfo(source, m)
}

func (sm *RoomSceneMgr) OnInit(mgr GameManager) {
	sm.manager = mgr
}

func (sm *RoomSceneMgr) OnClientConnected(id uint32, addr string) {
	fmt.Println("room client connected ", id, addr)
}

func (sm *RoomSceneMgr) OnClientDisconnected(id uint32) {
	fmt.Println("room client disconnected ", id)
}

func (sm *RoomSceneMgr) OnClientMessage(id uint32, m *proto.Message) {
	fmt.Println("room scene manager recv message .. ", id, m)
	sm.SendClientMessage(id, &proto.Message{
		Cmd: 101,
		Msg: &proto.GuestAuthMessage {
			Account: "hello world",
		},
	})
}

func (sm *RoomSceneMgr) OnCommandInfo(source uint32, m *proto.Message) {

}

func (sm *RoomSceneMgr) CreateScene(opt *SceneOption) Scene {
	return nil
}

