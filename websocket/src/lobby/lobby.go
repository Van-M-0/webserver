package lobby

import (
	"scenemgr"
	"gateway"
	"proto"
)

type LobbServer struct {
	mgr      		scenemgr.SceneManager
	gwServer 		*gateway.Gateway
}

func NewSingleLobbyServer(mgr scenemgr.SceneManager, gw *gateway.Gateway) *LobbServer {
	ls := &LobbServer{}

	gmgr := &LobbyManager{
		lsServer: ls,
		manager: mgr,
	}
	mgr.OnInit(gmgr)

	ls.gwServer = gw
	gw.RegisterNoitfier(&gateway.GateOption{
		Type: "lobby",
		Active: func(uid uint32, addr string) {
			gmgr.OnClientConnected(uid, addr)
		},
		Close: func(uid uint32) {
			gmgr.OnClientDisconnected(uid)
		},
		Auth: func(uid uint32, addr string) error {
			return nil
		},
		Msg: func(uid uint32, message *proto.Message) {
			gmgr.OnClientMessage(uid, message)
		},
	})

	return ls
}

func (ls *LobbServer) Start() {

}