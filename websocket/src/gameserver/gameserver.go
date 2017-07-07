package gameserver

import (
	"tcpserver"
	"export"
	"fmt"
	"proto"
)

type GameServer struct {
	gwClient 		*tcpserver.TcpClient
	mgr 			export.SceneManager
}

func NewGameServer(manager export.SceneManager) *GameServer {
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

	gf := &GameFrame{
		gameServer: gs,
	}

	gs.mgr = manager
	manager.OnInit(gf)

	gs.gwClient = cli

	return gs
}




