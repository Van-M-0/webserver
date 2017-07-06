package gameserver

import (
	"tcpserver"
	"proto"
	"fmt"
)

type GameServer struct {
	gwClient 		*tcpserver.TcpClient
}

func NewGameServer() *GameServer {
	gs := &GameServer{}

	cli, err := tcpserver.NewDailClient(&tcpserver.TcpOption{
		Addr: ":9092",
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
		fmt.Println("create gw client error")
		return nil
	}

	gs.gwClient = cli

	return gs
}

func (gs *GameServer) Start() {
	go gs.tcpServer.Start()
}


