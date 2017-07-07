package gameserver

import (
	"tcpserver"
	"scenemgr"
	"fmt"
	"proto"
	"gateway"
)

type GameServer struct {
	gwClient 		*tcpserver.TcpClient
	mgr      		scenemgr.SceneManager
	gwServer 		*gateway.Gateway
}

func NewSingleGameServer(mgr scenemgr.SceneManager, gw *gateway.Gateway) *GameServer {
	gs := &GameServer{}

	gmgr := &GameServerMgr{
		gameServer: gs,
		manager: mgr,
	}
	mgr.OnInit(gmgr)

	gs.gwServer = gw
	gw.RegisterNoitfier(&gateway.GateOption{
		Type: "gameserver",
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

	return gs
}

func newGameServer(manager scenemgr.SceneManager) *GameServer {
	gs := &GameServer{}
	cli, err := tcpserver.NewDailClient(&tcpserver.TcpOption{
		Addr: ":9099",
		Activecb: func(cli *tcpserver.TcpClient) {

		},
		Closecb: func(cli *tcpserver.TcpClient) {

		},
		Auth: func(cli *tcpserver.TcpClient) error {
			return nil
		},
		Msgcb: func(cli *tcpserver.TcpClient, m *proto.Message) {

		},
	})
	if err != nil {
		fmt.Println("create gw client error ", err)
		return nil
	}

	gf := &GameServerMgr{
		gameServer: gs,
	}

	gs.mgr = manager
	manager.OnInit(gf)

	gs.gwClient = cli

	return gs
}

func (gs *GameServer) Start() {
}



